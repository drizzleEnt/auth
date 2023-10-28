package service

import (
	"context"

	model1 "github.com/drizzleent/auth/internal/model"
)

type AuthService interface {
	Create(ctx context.Context, user *model1.User) (int, error)
	Get(ctx context.Context, id int64) (*model1.User, error)
	Update(ctx context.Context, user *model1.User) error
	Delete(ctx context.Context, id int64) error
}
