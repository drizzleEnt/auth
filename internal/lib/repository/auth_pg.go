package repository

import (
	"context"
	"fmt"

	"github.com/drizzleent/auth"
	"github.com/jackc/pgx/v4"
)

type AuthPostgres struct {
	db *pgx.Conn
}

func NewAuthPostgres(db *pgx.Conn) *AuthPostgres {
	return &AuthPostgres{
		db: db,
	}
}

func (r *AuthPostgres) CreateUser(ctx context.Context, user auth.User) (int, error) {

	var id int

	quary := fmt.Sprintf("INSERT INTO %s (name, email, role, password) values ($1, $2, $3, $4) RETURNING id", userTable)
	row := r.db.QueryRow(ctx, quary, user.Name, user.Email, user.Role, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(ctx context.Context, id int) (auth.User, error) {

	var user auth.User

	quary := fmt.Sprintf("SELECT name, email, password, role FROM %s WHERE id = $1", userTable)

	err := r.db.QueryRow(ctx, quary, id).Scan(&user.Name, &user.Email, &user.Password, &user.Role)

	return user, err

}
