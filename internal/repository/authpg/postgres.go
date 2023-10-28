package authpg

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

const (
	userTable        = "users"
	idColumn         = "id"
	nameColumn       = "name"
	emailColumn      = "email"
	roleColumn       = "role"
	passwordColumn   = "password"
	ctraetedAtColumn = "created_at"
	updatedAtColumn  = "updated_at"
)

type Config struct {
	Host     string
	Port     string
	UserName string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDb(cfg Config) (*pgx.Conn, error) {

	ctx := context.Background()
	con, err := pgx.Connect(ctx, fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.DBName, cfg.UserName, cfg.Password, cfg.SSLMode))

	if err != nil {
		return nil, err
	}

	//defer con.Close(ctx)

	err = con.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return con, nil
}
