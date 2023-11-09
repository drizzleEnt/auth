package authpg

import (
	"context"

	"github.com/drizzleent/auth/internal/client/db"
	"github.com/drizzleent/auth/internal/model"
	"github.com/drizzleent/auth/internal/repository"
)

type repo struct {
	db db.Client
}

func NewRepository(dbClient db.Client) repository.Authorisation {
	return &repo{db: dbClient}
}

func (r *repo) Create(ctx context.Context, user *model.User) (int, error) {

	// var id int

	// quary := fmt.Sprintf("INSERT INTO %s (%s, %s, %s, %s) values ($1, $2, $3, $4) RETURNING id", userTable, nameColumn, emailColumn, roleColumn, passwordColumn)
	// row := r.db.DB().QuaryRowContext(ctx, quary, user.Name, user.Email, user.Role, user.Password)

	// if err := row.Scan(&id); err != nil {
	// 	return 0, err
	// }

	// return id, nil
	return 0, nil
}

func (r *repo) Get(ctx context.Context, id int) (*model.User, error) {

	// var user data_model.User

	// quary := fmt.Sprintf("SELECT %s, %s, %s, %s FROM %s WHERE id = $1", nameColumn, emailColumn, passwordColumn, roleColumn, userTable)

	// err := r.db.QueryRow(ctx, quary, id).Scan(&user.Name, &user.Email, &user.Password, &user.Role)

	// return converter.ToModleUserFromRepo(&user), err

	return nil, nil

}

func (r *repo) Update(ctx context.Context, u *model.User) error {

	// quary := fmt.Sprintf("UPDATE %s SET %s=$1, %s=$2, %s=$3, %s=$4 WHERE %s=$5 RETURNING id", userTable, nameColumn, emailColumn, roleColumn, updatedAtColumn, idColumn)
	// _, err := r.db.Exec(ctx, quary, u.Name, u.Email, u.Role, time.Now())
	// return err
	return nil
}

func (r *repo) Delete(ctx context.Context, id int64) error {

	// quary := fmt.Sprintf("DELETE FROM %s WHERE %s=$1", userTable, idColumn)
	// _, err := r.db.Exec(ctx, quary, id)

	// return err

	return nil
}
