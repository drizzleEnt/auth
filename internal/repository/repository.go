package repository

import (
	"context"

	"github.com/drizzleent/auth/internal/model"
	"github.com/drizzleent/auth/internal/repository/authpg"
	"github.com/jackc/pgx/v4"
)

type Authorisation interface {
	Create(context.Context, *model.User) (int, error)
	Get(context.Context, int) (*model.User, error)
	Update(ctx context.Context, u *model.User) error
	Delete(ctx context.Context, id int64) error
}

type Ropository struct {
	Authorisation
}

func NewRepository(db *pgx.Conn) *Ropository {
	return &Ropository{
		Authorisation: authpg.NewAuthPostgres(db),
	}
}
