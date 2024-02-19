package access

import (
	"context"

	"github.com/drizzleent/auth/internal/client/db"
	"github.com/drizzleent/auth/internal/repository"
	"github.com/drizzleent/auth/internal/repository/access/model"
)

type repo struct {
	db db.Client
}

func NewAccessRepository(db db.Client) repository.AccessRepository {
	return &repo{
		db: db,
	}
}

func (r *repo) Check(ctx context.Context) (map[string]string, error) {
	accessbileRoles := make(map[string]string)
	accessbileRoles[model.ExamplePath] = "admin"

	return accessbileRoles, nil
}
