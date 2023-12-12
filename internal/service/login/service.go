package login

import (
	"github.com/drizzleent/auth/internal/repository"
	"github.com/drizzleent/auth/internal/service"
)

type serviceLogin struct {
	loginRepository repository.LoginRepository
}

func NewLoginService(loginRepository repository.LoginRepository) service.LoginService {
	return &serviceLogin{
		loginRepository: loginRepository,
	}
}
