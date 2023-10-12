package app

import (
	"database/sql"
	"time"
	"fmt"

	_ "github.com/lib/pq"
)

func NewDB(c *ConfigApplication) *sql.DB {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
						c.Database.User,
						c.Database.Password,
						c.Database.Host,
						c.Database.Port,
						c.Database.Name)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
