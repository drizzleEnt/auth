package app

import (
	"context"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/drizzleent/auth/internal/closer"
	"github.com/drizzleent/auth/internal/config"
	"github.com/drizzleent/auth/internal/interseptor"
	desc "github.com/drizzleent/auth/pkg/user_v2"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type App struct {
	serviceprovider *serviceProvider
	grpcServer      *grpc.Server
	httpServer      *http.Server
}

func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	wg := sync.WaitGroup{}

	wg.Add(2)

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
	a.grpcServer = grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		grpc.UnaryInterceptor(interseptor.ValidateInterceptor),
	)
	reflection.Register(a.grpcServer)

	desc.RegisterUserV1Server(a.grpcServer, a.serviceprovider.AuthImpl(ctx))
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

	log.Printf("HTTP server is running on %s", a.serviceprovider.httpConfig.Address())

	err := a.httpServer.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
