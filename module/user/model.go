package user

import (
	"database/sql"
	"context"
	"time"
)

type UserDatabaseIO struct {
	Id        int
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
	IsActive  sql.NullBool
}
