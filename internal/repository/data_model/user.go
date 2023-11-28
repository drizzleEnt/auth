package data_model

import (
	"database/sql"
	"time"
)

type UserUpdate struct {
	ID    int64  `db:"id"`
	Name  string `db:"username"`
	Email string `db:"email"`
	Role  int    `db:"role"`
}

type UserCreate struct {
	UserUpdate UserUpdate `db:""`
	Password   string     `db:"password"`
}

type User struct {
	UserCreate UserCreate   `db:""`
	CreatedAt  time.Time    `db:"created_at"`
	UpdatedAt  sql.NullTime `db:"updated_at"`
}
