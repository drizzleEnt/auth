package env

import (
	"fmt"
	"os"

	"github.com/drizzleent/auth/internal/config"
	"github.com/pkg/errors"
)

var _ config.PGConfig = (*pgConfig)(nil)

type pgConfig struct {
	dsn string
}

const (
	dbname     = "PG_DATABASE_NAME"
	dbuser     = "PG_USER"
	dbpassword = "PG_PASSWORD"
	dbport     = "PG_PORT_AUTH"
	dbhost     = "PG_HOST"
	dbssl      = "PG_SSL"
)

func NewPGConfig() (*pgConfig, error) {
	host := os.Getenv(dbhost)

	if len(host) == 0 {
		return nil, errors.New("db host not found")
	}

	port := os.Getenv(dbport)

	if len(port) == 0 {
		return nil, errors.New("db port not found")
	}

	name := os.Getenv(dbname)

	if len(name) == 0 {
		return nil, errors.New("db name not found")
	}

	user := os.Getenv(dbuser)

	if len(user) == 0 {
		return nil, errors.New("db user not found")
	}

	password := os.Getenv(dbpassword)

	if len(password) == 0 {
		return nil, errors.New("db password not found")
	}

	ssl := os.Getenv(dbssl)

	if len(ssl) == 0 {
		return nil, errors.New("db ssl not found")
	}

	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		host, port, name, user, password, ssl)

	return &pgConfig{
		dsn: dsn,
	}, nil
}

func (cfg *pgConfig) DSN() string {
	return cfg.dsn
}
