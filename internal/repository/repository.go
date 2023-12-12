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

type AccessRepository interface {
	Check(ctx context.Context) (map[string]string, error)
}

type LoginRepository interface {
	Login(ctx context.Context, info *model.UserClaims) (string, error)
	GetAccessToken(ctx context.Context, token string) (string, error)
	GetRefreshToken(ctx context.Context, tokrn string) (string, error)
	GetUserRole(ctx context.Context) (string, error)
}
