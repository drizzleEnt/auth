package auth

import (
	"context"
	"errors"

	"github.com/drizzleent/auth/internal/model"
)

func (s *Service) Get(ctx context.Context, id int64) (*model.User, error) {

	u, err := s.repo.Get(ctx, id)

	if err != nil {
		return nil, errors.New("Service.Get: " + err.Error())
	}

	return u, nil
}
