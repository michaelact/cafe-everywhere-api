package menu

import (
	"database/sql"
	"time"
)

type MenuDatabaseIO struct {
	Id          int
	Title       string
	Description string
	Count       int
	Price       int
	CafeId      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime
	IsActive    bool
}

type HTTPMenuRequest struct {
	Title       string `validate:"required" json:"title"`	
	CafeId      int    `validate:"required" json:"cafeId"`
	Count       int    `validate:"required" json:"count"`
	Price       int    `validate:"required" json:"price"`
	Description string `json:"description"`
}

type HTTPMenuUpdateRequest struct {
	Id           int    `validate:"required" json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Count        int    `json:"count"`
	Price        int    `json:"price"`
	CafeId       int    `json:"cafeId"`	
}

type HTTPMenuResponse struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Count       int       `json:"count"`
	Price       int       `json:"price"`
	CafeId      int       `json:"cafeId"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
