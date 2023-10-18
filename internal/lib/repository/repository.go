package repository

import (
	"context"

	"github.com/drizzleent/auth"
	"github.com/jackc/pgx/v4"
)

type Authorisation interface {
	CreateUser(context.Context, auth.User) (int, error)
	GetUser(context.Context, int) (auth.User, error)
}

type Ropository struct {
	Authorisation
}

func NewRepository(db *pgx.Conn) *Ropository {
	return &Ropository{
		Authorisation: NewAuthPostgres(db),
	}
}
