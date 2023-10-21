package user

import (
	"database/sql"
	"time"
)

type UserDatabaseIO struct {
	Id        int
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
	IsActive  bool
}

type HTTPUserRequest struct {
	Email     string `validate:"required" json:"email"`
	Password  string `validate:"required" json:"password"`
}

type HTTPUserUpdateRequest struct {
	Id        int    `validate:"required" json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type HTTPUserResponse struct {
	Id        int        `json:"id"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
