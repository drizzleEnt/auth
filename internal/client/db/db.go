package db

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type Handler func(ctx context.Context) error

type Client interface {
	DB() DB
	Close() error
}

type SQLExecer interface {
	NamedExecer
	QuaryExecer
}

type NamedExecer interface {
	ScanOneContext(context.Context, interface{}, Quary, ...interface{}) error
	ScanAllContext(context.Context, interface{}, Quary, ...interface{}) error
}

type QuaryExecer interface {
	ExecContext(context.Context, Quary, ...interface{}) (pgconn.CommandTag, error)
	QuaryContext(context.Context, Quary, ...interface{}) (pgx.Rows, error)
	QuaryRowContext(context.Context, Quary, ...interface{}) pgx.Row
}

type Transactor interface {
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
}

type Pinger interface {
	Ping(ctx context.Context) error
}

type Quary struct {
	Name     string
	QuaryRow string
}

type DB interface {
	SQLExecer
	Transactor
	Pinger
	Close()
}

type TxManager interface {
	ReadCommitted(context.Context, Handler) error
}
