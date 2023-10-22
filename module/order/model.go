package order

import (
	"database/sql"
	"time"
)

type OrderDatabaseIO struct {
	Id        int
	Notes     string
	Count     int
	MenuId    int
	UserId    int
	Status    string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
	IsActive  bool
}

type HTTPOrderRequest struct {
	MenuId  int    `validate:"required" json:"menu_id"`
	UserId  int    `validate:"required" json:"user_id"`
	Count   int    `validate:"required" json:"count"`
	Address string `validate:"required" json:"address"`
	Notes   string `json:"notes"`
}

type HTTPOrderUpdateRequest struct {
	Id     int    `validate:"required" json:"id"`
	Status string `json:"status"`
}

type HTTPOrderResponse struct {
	Id          int       `json:"id"`
	Notes       string    `json:"notes"`
	Count       int       `json:"count"`
	MenuId      int       `json:"menu_id"`
	UserId      int       `json:"user_id"`
	Status      string    `json:"status"`
	Address     string    `json:"address"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
