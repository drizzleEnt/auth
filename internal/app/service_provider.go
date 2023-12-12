package app

import (
	"context"
	"log"

	"github.com/drizzleent/auth/internal/api/access"
	"github.com/drizzleent/auth/internal/api/auth"
	"github.com/drizzleent/auth/internal/api/login"
	"github.com/drizzleent/auth/internal/client/db"
	"github.com/drizzleent/auth/internal/client/db/pg"
	"github.com/drizzleent/auth/internal/config"
	"github.com/drizzleent/auth/internal/config/env"
	"github.com/drizzleent/auth/internal/repository"
	accessRepository "github.com/drizzleent/auth/internal/repository/access"
	authRepository "github.com/drizzleent/auth/internal/repository/authpg"
	loginRepository "github.com/drizzleent/auth/internal/repository/login"
	"github.com/drizzleent/auth/internal/service"
	accessService "github.com/drizzleent/auth/internal/service/access"
	authService "github.com/drizzleent/auth/internal/service/auth"
	loginService "github.com/drizzleent/auth/internal/service/login"
)

type serviceProvider struct {
	pgConfig     config.PGConfig
	grpcCofig    config.GRPCConfig
	httpConfig   config.HTTPConfig
	swaggerCofig config.SwaggerConfig

	dbClient db.Client
	//txManager      db.TxManager

	authRepository   repository.Authorisation
	accessRepository repository.AccessRepository
	loginRepository  repository.LoginRepository

	authService   service.AuthService
	accessService service.AccessService
	loginService  service.LoginService

	authImp  *auth.Implementation
	accesImp *access.Implementation
	loginImp *login.Implementation
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

func (s *serviceProvider) SwaggerConfig() config.SwaggerConfig {
	if s.swaggerCofig == nil {
		cfg, err := env.NewSwaggerConfig()
		if err != nil {
			log.Fatalf("failed to load swagger config: %s", err.Error())
		}
		s.swaggerCofig = cfg
	}

	return s.swaggerCofig
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

func (s *serviceProvider) AccessRepository(ctx context.Context) repository.AccessRepository {
	if s.accessRepository == nil {
		s.accessRepository = accessRepository.NewAccessRepository(s.DbClient(ctx))
	}
	return s.accessRepository
}

func (s *serviceProvider) LoginRepository(ctx context.Context) repository.LoginRepository {
	if s.loginRepository == nil {
		s.loginRepository = loginRepository.NewLoginRepository(s.DbClient(ctx))
	}

	return s.loginRepository
}

func (s *serviceProvider) AuthService(ctx context.Context) service.AuthService {
	if s.authService == nil {
		s.authService = authService.NewService(s.AuthRepository(ctx))
	}

	return s.authService
}

func (s *serviceProvider) AccesService(ctx context.Context) service.AccessService {
	if s.accessService == nil {
		s.accessService = accessService.NewAccessService(s.AccessRepository(ctx))
	}

	return s.accessService
}

func (s *serviceProvider) LoginService(ctx context.Context) service.LoginService {
	if s.loginService == nil {
		s.loginService = loginService.NewLoginService(s.LoginRepository(ctx))
	}

	return s.loginService
}

func (s *serviceProvider) AuthImpl(ctx context.Context) *auth.Implementation {
	if s.authImp == nil {
		s.authImp = auth.NewImplementation(s.AuthService(ctx))
	}
	return s.authImp
}

func (s *serviceProvider) AccessImpl(ctx context.Context) *access.Implementation {
	if s.accesImp == nil {
		s.accesImp = access.NewImplementation(s.AccesService(ctx))
	}

	return s.accesImp
}

func (s *serviceProvider) LoginImpl(ctx context.Context) *login.Implementation {
	if s.loginImp == nil {
		s.loginImp = login.NewImplementation(s.LoginService(ctx))
	}

	return s.loginImp
}
