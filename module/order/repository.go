package order

import (
	"database/sql"
	"context"
	"errors"

	"github.com/michaelact/cafe-everywhere/helper"
)

type OrderRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, order OrderDatabaseIO) OrderDatabaseIO
	Update(ctx context.Context, tx *sql.Tx, order OrderDatabaseIO) OrderDatabaseIO
	Delete(ctx context.Context, tx *sql.Tx, id int)
	FindById(ctx context.Context, tx *sql.Tx, id int) (OrderDatabaseIO, error)
	FindByUserId(ctx context.Context, tx *sql.Tx, userId int) []OrderDatabaseIO
	FindAll(ctx context.Context, tx *sql.Tx) []OrderDatabaseIO
}

type OrderRepositoryImpl struct {}

func NewOrderRepository() OrderRepository {
	return &OrderRepositoryImpl{}
}

func (self *OrderRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, order OrderDatabaseIO) OrderDatabaseIO {
	// Insert new order
	SQLPut := "INSERT INTO \"order\"(menu_id, user_id, notes, count, address) VALUES ($1,$2,$3,$4,$5) RETURNING id,status,created_at,updated_at;"
	tx.QueryRowContext(ctx, SQLPut, order.MenuId, order.UserId, order.Notes, order.Count, order.Address).Scan(&order.Id, &order.Status, &order.CreatedAt, &order.UpdatedAt)

	SQLReduceAvailableMenu := "UPDATE menu SET count = count - $1 WHERE id=$2;"
	_, err := tx.ExecContext(ctx, SQLReduceAvailableMenu, order.Count, order.MenuId)
	helper.PanicIfError(err)

	// Return created order
	return order
}

func (self *OrderRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, order OrderDatabaseIO) OrderDatabaseIO {
	// Update existing order
	SQLPut := "UPDATE \"order\" SET status=$1, updated_at=NOW() WHERE id=$2 AND is_active=true;"
	_, err := tx.ExecContext(ctx, SQLPut, order.Status, order.Id)
	helper.PanicIfError(err)

	// Return updated order
	order, err = self.FindById(ctx, tx, order.Id)
	helper.PanicIfError(err)
	return order
}

func (self *OrderRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) {
	// Delete existing order
	SQLDel := "UPDATE \"order\" SET deleted_at=NOW(), is_active=false WHERE id=$1;"
	_, err := tx.ExecContext(ctx, SQLDel, id)
	helper.PanicIfError(err)

	SQLRefundAvailableCount := "UPDATE menu SET count = count + ( SELECT SUM(count) FROM \"order\" WHERE id = $1 ) WHERE id = ( SELECT menu_id FROM \"order\" WHERE id = $1 );"
	_, err = tx.ExecContext(ctx, SQLRefundAvailableCount, id)
	helper.PanicIfError(err)
}

func (self *OrderRepositoryImpl) FindByUserId(ctx context.Context, tx *sql.Tx, userId int) []OrderDatabaseIO {
	// Extract existing order
	SQLGet := "SELECT id, menu_id, notes, count, status, address, created_at, updated_at FROM \"order\" WHERE user_id=$1 AND is_active=true;"
	rows, err := tx.QueryContext(ctx, SQLGet, userId)
	helper.PanicIfError(err)

	// Iterate all extracted rows
	var listOrder []OrderDatabaseIO
	defer rows.Close()
	for rows.Next() {
		order := OrderDatabaseIO{}
		order.UserId = userId
		err := rows.Scan(&order.Id, &order.MenuId, &order.Notes, &order.Count, &order.Status, &order.Address, &order.CreatedAt, &order.UpdatedAt)
		helper.PanicIfError(err)

		listOrder = append(listOrder, order)
	}

	return listOrder
}

func (self *OrderRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (OrderDatabaseIO, error) {
	// Extract existing order
	SQLGet := "SELECT menu_id, user_id, notes, count, status, address, created_at, updated_at FROM \"order\" WHERE id=$1 AND is_active=true LIMIT 1;"
	rows, err := tx.QueryContext(ctx, SQLGet, id)
	helper.PanicIfError(err)

	// Bind all columns value to order variable
	order := OrderDatabaseIO{}
	order.Id = id
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&order.MenuId, &order.UserId, &order.Notes, &order.Count, &order.Status, &order.Address, &order.CreatedAt, &order.UpdatedAt)
		helper.PanicIfError(err)
		return order, nil
	} else {
		return order, errors.New("Order not found")
	}
}

func (self *OrderRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []OrderDatabaseIO {
	// Extract all order
	SQLGet := "SELECT id, menu_id, user_id, notes, count, status, address, created_at, updated_at FROM \"order\" WHERE is_active=true"
	rows, err := tx.QueryContext(ctx, SQLGet)
	helper.PanicIfError(err)

	// Iterate all extracted rows
	var listOrder []OrderDatabaseIO
	defer rows.Close()
	for rows.Next() {
		order := OrderDatabaseIO{}
		err := rows.Scan(&order.Id, &order.MenuId, &order.UserId, &order.Notes, &order.Count, &order.Status, &order.Address, &order.CreatedAt, &order.UpdatedAt)
		helper.PanicIfError(err)

		listOrder = append(listOrder, order)
	}

	return listOrder
}
