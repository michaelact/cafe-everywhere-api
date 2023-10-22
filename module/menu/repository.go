package menu

import (
	"database/sql"
	"context"
	"errors"

	"github.com/michaelact/cafe-everywhere/helper"
)

type MenuRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, menu MenuDatabaseIO) MenuDatabaseIO
	Update(ctx context.Context, tx *sql.Tx, menu MenuDatabaseIO) MenuDatabaseIO
	Delete(ctx context.Context, tx *sql.Tx, id int)
	FindById(ctx context.Context, tx *sql.Tx, id int) (MenuDatabaseIO, error)
	FindByCafeId(ctx context.Context, tx *sql.Tx, cafeId int) []MenuDatabaseIO
	FindAll(ctx context.Context, tx *sql.Tx) []MenuDatabaseIO
}

type MenuRepositoryImpl struct {}

func NewMenuRepository() MenuRepository {
	return &MenuRepositoryImpl{}
}

func (self *MenuRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, menu MenuDatabaseIO) MenuDatabaseIO {
	// Insert new menu
	SQLPut := "INSERT INTO menu(cafe_id, title, description, count, price) VALUES ($1,$2,$3,$4,$5) RETURNING id;"
	tx.QueryRowContext(ctx, SQLPut, menu.CafeId, menu.Title, menu.Description, menu.Count, menu.Price).Scan(&menu.Id)

	// Return created menu
	return menu
}

func (self *MenuRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, menu MenuDatabaseIO) MenuDatabaseIO {
	// Update existing menu
	SQLPut := "UPDATE menu SET cafe_id=$1, title=$2, price=$3, description=$4, count=$5, price=$6, updated_at=NOW() WHERE id=$7 AND is_active=true;"
	_, err := tx.ExecContext(ctx, SQLPut, menu.CafeId, menu.Title, menu.Price, menu.Description, menu.Count, menu.Price, menu.Id)
	helper.PanicIfError(err)

	// Return updated menu
	menu, err = self.FindById(ctx, tx, menu.Id)
	helper.PanicIfError(err)
	return menu
}

func (self *MenuRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) {
	// Delete existing menu
	SQLDel := "UPDATE menu SET deleted_at=NOW(), is_active=false WHERE id=$1;"
	_, err := tx.ExecContext(ctx, SQLDel, id)
	helper.PanicIfError(err)
}

func (self *MenuRepositoryImpl) FindByCafeId(ctx context.Context, tx *sql.Tx, cafeId int) []MenuDatabaseIO {
	// Extract existing menu
	SQLGet := "SELECT id, title, description, count, price, created_at, updated_at FROM menu WHERE cafe_id=$1 AND is_active=true;"
	rows, err := tx.QueryContext(ctx, SQLGet, cafeId)
	helper.PanicIfError(err)

	// Iterate all extracted rows
	var listMenu []MenuDatabaseIO
	defer rows.Close()
	for rows.Next() {
		menu := MenuDatabaseIO{}
		menu.CafeId = cafeId
		err := rows.Scan(&menu.Id, &menu.Title, &menu.Description, &menu.Count, &menu.Price, &menu.CreatedAt, &menu.UpdatedAt)
		helper.PanicIfError(err)

		listMenu = append(listMenu, menu)
	}

	return listMenu
}

func (self *MenuRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (MenuDatabaseIO, error) {
	// Extract existing menu
	SQLGet := "SELECT cafe_id, title, description, count, price, created_at, updated_at FROM menu WHERE id=$1 AND is_active=true LIMIT 1;"
	rows, err := tx.QueryContext(ctx, SQLGet, id)
	helper.PanicIfError(err)

	// Bind all columns value to menu variable
	menu := MenuDatabaseIO{}
	menu.Id = id
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&menu.CafeId, &menu.Title, &menu.Description, &menu.Count, &menu.Price, &menu.CreatedAt, &menu.UpdatedAt)
		helper.PanicIfError(err)
		return menu, nil
	} else {
		return menu, errors.New("Menu not found")
	}
}

func (self *MenuRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []MenuDatabaseIO {
	// Extract all menu
	SQLGet := "SELECT id, cafe_id, title, description, count, price, created_at, updated_at FROM menu WHERE is_active=true"
	rows, err := tx.QueryContext(ctx, SQLGet)
	helper.PanicIfError(err)

	// Iterate all extracted rows
	var listMenu []MenuDatabaseIO
	defer rows.Close()
	for rows.Next() {
		menu := MenuDatabaseIO{}
		err := rows.Scan(&menu.Id, &menu.CafeId, &menu.Title, &menu.Description, &menu.Count, &menu.Price, &menu.CreatedAt, &menu.UpdatedAt)
		helper.PanicIfError(err)

		listMenu = append(listMenu, menu)
	}

	return listMenu
}
