package model

import (
	"database/sql"
	"time"
)

type User struct {
	UserCreate UserCreate   `db:""`
	CreatedAt  time.Time    `db:"created_at"`
	UpdatedAt  sql.NullTime `db:"updated_at"`
}

type UserCreate struct {
	UserUpdate UserUpdate `db:""`
	Password   string     `db:"password"`
}

type UserUpdate struct {
	ID    int64  `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
	Role  int    `db:"role"`
}
