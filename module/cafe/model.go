package cafe

import (
	"database/sql"
	"time"
)

type CafeDatabaseIO struct {
	Id        int
	Email     string
	Title     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
	IsActive  bool
}

type HTTPCafeRequest struct {
	Email     string `validate:"required" json:"email"`
	Title     string `json:"title"`
	Password  string `json:"password"`
}

type HTTPCafeUpdateRequest struct {
	Id        int    `validate:"required" json:"id"`
	Email     string `json:"email"`
	Title     string `json:"title"`
	Password  string `json:"password"`
}

type HTTPCafeResponse struct {
	Id        int        `json:"id"`
	Email     string     `json:"email"`
	Title     string     `json:"title"`
	Password  string     `json:"password"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
