package login

import (
	"github.com/drizzleent/auth/internal/service"
	desc "github.com/drizzleent/auth/pkg/login_v1"
)

type Implementation struct {
	desc.UnimplementedLoginV1Server
	loginService service.LoginService
}

func NewImplementation(loginService service.LoginService) *Implementation {
	return &Implementation{
		loginService: loginService,
	}
}
