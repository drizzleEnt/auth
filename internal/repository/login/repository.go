package login

import (
	"context"

	"github.com/drizzleent/auth/internal/client/db"
	"github.com/drizzleent/auth/internal/model"
	"github.com/drizzleent/auth/internal/repository"
)

type repo struct {
	db db.Client
}

func NewLoginRepository(db db.Client) repository.LoginRepository {
	return &repo{
		db: db,
	}
}

func (r *repo) Login(ctx context.Context, info *model.UserClaims) (string, error) {
	return "", nil
}
func (r *repo) GetAccessToken(ctx context.Context, token string) (string, error) {
	return "", nil
}
func (r *repo) GetRefreshToken(ctx context.Context, tokrn string) (string, error) {
	return "", nil
}
func (r *repo) GetUserRole(ctx context.Context) (string, error) {
	return "", nil
}
