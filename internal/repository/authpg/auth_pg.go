package authpg

import (
	"context"
	"fmt"
	"time"

	"github.com/drizzleent/auth/internal/model"
	"github.com/drizzleent/auth/internal/repository/converter"
	"github.com/drizzleent/auth/internal/repository/data_model"
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

func (r *AuthPostgres) Create(ctx context.Context, user *model.User) (int, error) {

	var id int

	quary := fmt.Sprintf("INSERT INTO %s (%s, %s, %s, %s) values ($1, $2, $3, $4) RETURNING id", userTable, nameColumn, emailColumn, roleColumn, passwordColumn)
	row := r.db.QueryRow(ctx, quary, user.Name, user.Email, user.Role, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) Get(ctx context.Context, id int) (*model.User, error) {

	var user data_model.User

	quary := fmt.Sprintf("SELECT %s, %s, %s, %s FROM %s WHERE id = $1", nameColumn, emailColumn, passwordColumn, roleColumn, userTable)

	err := r.db.QueryRow(ctx, quary, id).Scan(&user.Name, &user.Email, &user.Password, &user.Role)

	return converter.ToModleUserFromRepo(&user), err

}

func (r *AuthPostgres) Update(ctx context.Context, u *model.User) error {

	quary := fmt.Sprintf("UPDATE %s SET %s=$1, %s=$2, %s=$3, %s=$4 WHERE %s=$5 RETURNING id", userTable, nameColumn, emailColumn, roleColumn, updatedAtColumn, idColumn)
	_, err := r.db.Exec(ctx, quary, u.Name, u.Email, u.Role, time.Now())
	return err
}

func (r *AuthPostgres) Delete(ctx context.Context, id int64) error {

	quary := fmt.Sprintf("DELETE FROM %s WHERE %s=$1", userTable, idColumn)
	_, err := r.db.Exec(ctx, quary, id)

	return err
}
