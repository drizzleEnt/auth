package auth

import (
	"context"
	"errors"

	model1 "github.com/drizzleent/auth/internal/model"
)

func (s *Service) Create(ctx context.Context, user *model1.User) (int, error) {

	id, err := s.repo.Create(ctx, user)

	if err != nil {
		return 0, errors.New("Service.Create: " + err.Error())
	}

	return id, nil
}
