package auth

import (
	"context"

	"github.com/drizzleent/auth/internal/model"
)

func (s *Service) Create(ctx context.Context, user *model.UserCreate) (int64, error) {

	id, err := s.repo.Create(ctx, user)

	if err != nil {
		return 0, err
	}

	return id, nil
}
