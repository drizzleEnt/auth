package auth

import (
	"log"

	"github.com/drizzleent/auth/internal/service"
	desc "github.com/drizzleent/auth/pkg/user_v1"
)

type Implementation struct {
	desc.UnimplementedUserV1Server
	log         *log.Logger
	authservice service.AuthService
}

func NewUserRpcsServer(log *log.Logger, service service.AuthService) *Implementation {
	return &Implementation{
		log:         log,
		authservice: service,
	}
}
