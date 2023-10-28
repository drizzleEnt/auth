package auth

import (
	"context"
	"errors"

	model1 "github.com/drizzleent/auth/internal/model"
)

func (s *Service) Update(ctx context.Context, user *model1.User) error {
	err := s.repo.Update(ctx, user)
	if err != nil {
		return errors.New("Service.Update: " + err.Error())
	}
	return nil
}
