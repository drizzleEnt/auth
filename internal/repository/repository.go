package repository

import (
	"context"

	"github.com/drizzleent/auth/internal/model"
)

type Authorisation interface {
	Create(ctx context.Context, user *model.UserCreate) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
	Update(ctx context.Context, u *model.UserUpdate) error
	Delete(ctx context.Context, id int64) error
}
