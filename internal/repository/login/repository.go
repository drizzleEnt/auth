package login

import (
	"context"
	"fmt"

	"github.com/drizzleent/auth/internal/client/db"
	"github.com/drizzleent/auth/internal/model"
	"github.com/drizzleent/auth/internal/repository"
)

const (
	userTable      = "users"
	nameColumn     = "name"
	roleColumn     = "role"
	passwordColumn = "password"
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
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s=$1 AND %s=$2", roleColumn, userTable, nameColumn, passwordColumn)
	q := db.Quary{
		Name:     "repository.Login",
		QuaryRow: query,
	}

	args := []interface{}{info.Username, info.Password}

	var role int
	err := r.db.DB().QuaryRowContext(ctx, q, args...).Scan(&role)
	fmt.Println("role: ", role)
	return fmt.Sprint(role), err
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
