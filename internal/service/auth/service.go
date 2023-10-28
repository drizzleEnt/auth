package auth

import (
	"github.com/drizzleent/auth/internal/repository"
	"github.com/drizzleent/auth/internal/service"
)

type Service struct {
	repo repository.Authorisation
}

func NewService(repo repository.Authorisation) service.AuthService {
	return &Service{
		repo: repo,
	}
}
