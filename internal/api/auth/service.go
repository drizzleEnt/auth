package auth

import (
	"github.com/drizzleent/auth/internal/service"
	desc "github.com/drizzleent/auth/pkg/user_v2"
)

type Implementation struct {
	desc.UnimplementedUserV1Server
	authservice service.AuthService
}

func NewImplementation(service service.AuthService) *Implementation {

	return &Implementation{
		authservice: service,
	}
}
