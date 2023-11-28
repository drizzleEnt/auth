package auth

import (
	"context"
	"errors"

	"github.com/drizzleent/auth/internal/model"
)

func (s *Service) Create(ctx context.Context, user *model.UserCreate) (int64, error) {

	id, err := s.repo.Create(ctx, user)

	if err != nil {
		return 0, errors.New("Service.Create: " + err.Error())
	}

	return id, nil
}
