package auth

import (
	"context"
	"errors"

	"github.com/drizzleent/auth/internal/model"
)

func (s *Service) Update(ctx context.Context, user *model.UserUpdate) error {
	err := s.repo.Update(ctx, user)
	if err != nil {
		return errors.New("Service.Update: " + err.Error())
	}
	return nil
}
