package auth

import (
	"context"
	"errors"
)

func (s *Service) Delete(ctx context.Context, id int64) error {

	err := s.repo.Delete(ctx, id)
	if err != nil {
		return errors.New("Service:Delete: " + err.Error())
	}
	return nil
}
