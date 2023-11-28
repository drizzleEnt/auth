package authpg

import (
	"context"
	"fmt"
	"time"

	"github.com/drizzleent/auth/internal/client/db"
	"github.com/drizzleent/auth/internal/model"
	"github.com/drizzleent/auth/internal/repository"
	"github.com/drizzleent/auth/internal/repository/converter"
	"github.com/drizzleent/auth/internal/repository/data_model"
)

type repo struct {
	db db.Client
}

func NewRepository(dbClient db.Client) repository.Authorisation {
	return &repo{db: dbClient}
}

func (r *repo) Create(ctx context.Context, user *model.UserCreate) (int64, error) {

	quary := fmt.Sprintf("INSERT INTO %s (%s, %s, %s, %s) values ($1, $2, $3, $4) RETURNING id", userTable, nameColumn, emailColumn, roleColumn, passwordColumn)

	q := db.Quary{
		Name:     "auth.repository.Create",
		QuaryRow: quary,
	}

	args := []interface{}{user.UserUpdate.Name, user.UserUpdate.Email, user.UserUpdate.Role, user.Password}

	var id int64

	err := r.db.DB().QuaryRowContext(ctx, q, args...).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) Get(ctx context.Context, id int64) (*model.User, error) {

	var user data_model.User

	quary := fmt.Sprintf("SELECT %s, %s, %s, %s FROM %s WHERE id = $1", nameColumn, emailColumn, passwordColumn, roleColumn, userTable)

	q := db.Quary{
		Name:     "Create",
		QuaryRow: quary,
	}

	args := []interface{}{id}
	err := r.db.DB().QuaryRowContext(ctx, q, args...).Scan(&user.UserCreate.UserUpdate.Name,
		&user.UserCreate.UserUpdate.Email,
		&user.UserCreate.Password,
		&user.UserCreate.UserUpdate.Role)

	return converter.ToModleUserFromRepo(user), err

}

func (r *repo) Update(ctx context.Context, u *model.UserUpdate) error {

	quary := fmt.Sprintf("UPDATE %s SET %s=$1, %s=$2, %s=$3, %s=$4 WHERE %s=$5 RETURNING id", userTable, nameColumn, emailColumn, roleColumn, updatedAtColumn, idColumn)
	q := db.Quary{
		Name:     "auth.repository.Update",
		QuaryRow: quary,
	}
	args := []interface{}{u.Name, u.Email, u.Role, time.Now(), u.ID}
	res, err := r.db.DB().ExecContext(ctx, q, args...)

	if err != nil {
		return fmt.Errorf("Failed to Update user: %v, tag: %v", err, res)
	}

	return nil
}

func (r *repo) Delete(ctx context.Context, id int64) error {

	quary := fmt.Sprintf("DELETE FROM %s WHERE %s=$1", userTable, idColumn)
	q := db.Quary{
		Name:     "auth.repository.Delete",
		QuaryRow: quary,
	}
	args := []interface{}{id}
	res, err := r.db.DB().ExecContext(ctx, q, args...)

	if err != nil {
		return fmt.Errorf("Failed to Delete user: %v, tag: %v", err, res)
	}

	return nil
}
