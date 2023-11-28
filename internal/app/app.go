package app

import (
	"context"
	"log"
	"net"

	"github.com/drizzleent/auth/internal/closer"
	"github.com/drizzleent/auth/internal/config"
	desc "github.com/drizzleent/auth/pkg/user_v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type App struct {
	serviceprovider *serviceProvider
	grpcServer      *grpc.Server
}

func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	return a.runGrpcServer()
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
	a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	reflection.Register(a.grpcServer)

	desc.RegisterUserV1Server(a.grpcServer, a.serviceprovider.AuthImpl(ctx))
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
