package app

import (
	"context"
	"log"

	"github.com/drizzleent/auth/internal/api/auth"
	"github.com/drizzleent/auth/internal/client/db"
	"github.com/drizzleent/auth/internal/client/db/pg"
	"github.com/drizzleent/auth/internal/config"
	"github.com/drizzleent/auth/internal/config/env"
	"github.com/drizzleent/auth/internal/repository"
	authRepository "github.com/drizzleent/auth/internal/repository/authpg"
	"github.com/drizzleent/auth/internal/service"
	authService "github.com/drizzleent/auth/internal/service/auth"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcCofig  config.GRPCConfig
	httpConfig config.HTTPConfig

	dbClient db.Client
	//txManager      db.TxManager

	authRepository repository.Authorisation
	authService    service.AuthService
	authImp        *auth.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcCofig == nil {
		cfg, err := env.NewGrpcConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcCofig = cfg
	}

	return s.grpcCofig
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := env.NewHttpConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) DbClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("Failed to create db client %s", err.Error())
		}

		err = cl.DB().Ping(ctx)

		if err != nil {
			log.Fatalf("Failed to ping db %s", err.Error())
		}

		s.dbClient = cl
	}

	return s.dbClient
}

// func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager{
// 	 if s.txManager == nil{
// 		s.txManager =
// 	 }
// }

func (s *serviceProvider) AuthRepository(ctx context.Context) repository.Authorisation {
	if s.authRepository == nil {
		s.authRepository = authRepository.NewRepository(s.DbClient(ctx))
	}

	return s.authRepository
}

func (s *serviceProvider) AuthService(ctx context.Context) service.AuthService {
	if s.authService == nil {
		s.authService = authService.NewService(s.AuthRepository(ctx))
	}

	return s.authService
}

func (s *serviceProvider) AuthImpl(ctx context.Context) *auth.Implementation {
	if s.authImp == nil {
		s.authImp = auth.NewImplementation(s.AuthService(ctx))
	}
	return s.authImp
}
