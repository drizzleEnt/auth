package repository

import (
	"context"

	"github.com/drizzleent/auth/internal/model"
)

type Authorisation interface {
	Create(ctx context.Context, user *model.User) (int, error)
	Get(ctx context.Context, id int) (*model.User, error)
	Update(ctx context.Context, u *model.User) error
	Delete(ctx context.Context, id int64) error
}
