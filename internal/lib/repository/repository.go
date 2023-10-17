package repository

import (
	"github.com/drizzleent/auth"
	"github.com/jmoiron/sqlx"
)

type Authorisation interface {
	CreateUser(auth.User) (int, error)
	GetUser(int) (auth.User, error)
}

type Ropository struct {
	Authorisation
}

func NewRepository(db *sqlx.DB) *Ropository {
	return &Ropository{
		Authorisation: NewAuthPostgres(db),
	}
}
