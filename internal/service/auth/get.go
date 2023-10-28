package auth

import (
	"context"
	"errors"

	model1 "github.com/drizzleent/auth/internal/model"
)

func (s *Service) Get(ctx context.Context, id int64) (*model1.User, error) {

	u, err := s.repo.Get(ctx, int(id))
	if err != nil {
		return nil, errors.New("Service.Get: " + err.Error())
	}
	return u, nil
}
