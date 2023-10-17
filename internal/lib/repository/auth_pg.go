package repository

import (
	"fmt"

	"github.com/drizzleent/auth"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{
		db: db,
	}
}

func (r *AuthPostgres) CreateUser(user auth.User) (int, error) {

	var id int

	quary := fmt.Sprintf("INSERT INTO %s (name, email, role, password) values ($1, $2, $3, $4) RETURNING id", userTable)
	row := r.db.QueryRow(quary, user.Name, user.Email, user.Role, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(id int) (auth.User, error) {

	var user auth.User

	quary := fmt.Sprintf("SELECT name, email, password, role FROM %s WHERE id=$1", userTable)

	err := r.db.Get(&user, quary, id)

	return user, err
}
