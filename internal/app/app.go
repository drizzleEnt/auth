package app

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/natefinch/lumberjack"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rakyll/statik/fs"
	"github.com/rs/cors"

	"github.com/drizzleent/auth/internal/closer"
	"github.com/drizzleent/auth/internal/config"
	"github.com/drizzleent/auth/internal/interseptor"
	"github.com/drizzleent/auth/internal/logger"
	"github.com/drizzleent/auth/internal/metric"
	descAccess "github.com/drizzleent/auth/pkg/access_v1"
	descLogin "github.com/drizzleent/auth/pkg/login_v1"
	desc "github.com/drizzleent/auth/pkg/user_v2"
	_ "github.com/drizzleent/auth/statik"
)

var logLevel = flag.String("l", "info", "log level")

type App struct {
	serviceprovider  *serviceProvider
	grpcServer       *grpc.Server
	httpServer       *http.Server
	swaggerServer    *http.Server
	prometheusServer *http.Server
}

func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	wg := sync.WaitGroup{}

	wg.Add(4)

	go func() {
		defer wg.Done()

		err := a.runGrpcServer()
		if err != nil {
			log.Fatalf("Failed to run grpc server %s", err.Error())
		}
	}()

	go func() {
		defer wg.Done()

		err := a.runHTTPServer()
		if err != nil {
			log.Fatalf("Failed to run http server %s", err.Error())
		}
	}()

	go func() {
		defer wg.Done()

		err := a.runSwaggerServer()
		if err != nil {
			log.Fatalf("Failed to run swagger server %s", err.Error())
		}
	}()

	go func() {
		defer wg.Done()

		err := a.runPrometheusServer()
		if err != nil {
			log.Fatalf("Failed to run prometheus server %s", err.Error())
		}
	}()

	wg.Wait()

	return nil
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initGrpcServer,
		a.initHTTPServer,
		a.initSwaggerServer,
		a.initPrometheusServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load(".env")
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceprovider = newServiceProvider()
	return nil
}

func (a *App) initGrpcServer(ctx context.Context) error {
	logger.Init(getCore(getAtomicLevel()))
	err := metric.Init(ctx)
	if err != nil {
		return fmt.Errorf("failed to init metrics: %v", err)
	}

	a.grpcServer = grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		grpc.UnaryInterceptor(
			grpcMiddleware.ChainUnaryServer(
				interseptor.MetricInterceptor,
				interseptor.LogInterceptor,
				interseptor.ValidateInterceptor,
			),
		))
	reflection.Register(a.grpcServer)

	desc.RegisterUserV1Server(a.grpcServer, a.serviceprovider.AuthImpl(ctx))
	descAccess.RegisterAccessV1Server(a.grpcServer, a.serviceprovider.AccessImpl(ctx))
	descLogin.RegisterLoginV1Server(a.grpcServer, a.serviceprovider.LoginImpl(ctx))

	return nil
}

func (a *App) initHTTPServer(ctx context.Context) error {
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	err := desc.RegisterUserV1HandlerFromEndpoint(ctx, mux, a.serviceprovider.HTTPConfig().Address(), opts)

	if err != nil {
		return err
	}

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Authorization"},
		AllowCredentials: true,
	})

	a.httpServer = &http.Server{
		Addr:    a.serviceprovider.httpConfig.Address(),
		Handler: corsMiddleware.Handler(mux),
	}

	return nil
}

func (a *App) initSwaggerServer(ctx context.Context) error {

	statickFs, err := fs.New()
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.StripPrefix("/", http.FileServer(statickFs)))
	mux.HandleFunc("/api.swagger.json", serveSwaggerServer("/api.swagger.json"))

	a.swaggerServer = &http.Server{
		Addr:    a.serviceprovider.SwaggerConfig().Address(),
		Handler: mux,
	}

	return nil
}

func (a *App) initPrometheusServer(ctx context.Context) error {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	a.prometheusServer = &http.Server{
		Addr:    a.serviceprovider.PrometheusConfig().Address(),
		Handler: mux,
	}

	return nil
}

func (a *App) runGrpcServer() error {
	log.Printf("GRPC server is running on %s", a.serviceprovider.GRPCConfig().Address())

	list, err := net.Listen("tcp", a.serviceprovider.GRPCConfig().Address())

	if err != nil {
		return err
	}

	err = a.grpcServer.Serve(list)

	if err != nil {
		return err
	}

	return nil
}

func (a *App) runHTTPServer() error {

	log.Printf("HTTP server is running on %s", a.serviceprovider.HTTPConfig().Address())

	err := a.httpServer.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func (a *App) runSwaggerServer() error {
	log.Printf("Swagger server running on: %s", a.serviceprovider.SwaggerConfig().Address())

	err := a.swaggerServer.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func (a *App) runPrometheusServer() error {
	log.Printf("Prometheus server is running on %s", a.serviceprovider.PrometheusConfig().Address())

	err := a.prometheusServer.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func serveSwaggerServer(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Serving swagger file: %s", path)

		statickFs, err := fs.New()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Printf("Open swagger file: %s", path)

		file, err := statickFs.Open(path)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer file.Close()

		log.Printf("Read swagger file: %s", path)

		content, err := io.ReadAll(file)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-type", "application.json")
		_, err = w.Write(content)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Printf("Served swagger file: %s", path)

	}
}

func getCore(level zap.AtomicLevel) zapcore.Core {
	stdout := zapcore.AddSync(os.Stdout)

	file := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "logs/app.log",
		MaxSize:    10,
		MaxAge:     7,
		MaxBackups: 3,
	})

	productionCfg := zap.NewProductionEncoderConfig()
	productionCfg.TimeKey = "timestamp"
	productionCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	develomentCfg := zap.NewDevelopmentEncoderConfig()
	develomentCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

	consoleEncoder := zapcore.NewConsoleEncoder(develomentCfg)
	fileEncoder := zapcore.NewJSONEncoder(productionCfg)

	return zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, stdout, level),
		zapcore.NewCore(fileEncoder, file, level),
	)
}

func getAtomicLevel() zap.AtomicLevel {
	var level zapcore.Level

	if err := level.Set(*logLevel); err != nil {
		log.Fatalf("failed to set log level: %v", err)
	}

	return zap.NewAtomicLevelAt(level)
}
