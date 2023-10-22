package cafe

import (
	"database/sql"
	"context"
	"errors"

	"github.com/michaelact/cafe-everywhere/helper"
)

type CafeRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, cafe CafeDatabaseIO) CafeDatabaseIO
	Update(ctx context.Context, tx *sql.Tx, cafe CafeDatabaseIO) CafeDatabaseIO
	Delete(ctx context.Context, tx *sql.Tx, id int)
	FindById(ctx context.Context, tx *sql.Tx, id int) (CafeDatabaseIO, error)
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (CafeDatabaseIO, error)
	FindAll(ctx context.Context, tx *sql.Tx) []CafeDatabaseIO
}

type CafeRepositoryImpl struct {}

func NewCafeRepository() CafeRepository {
	return &CafeRepositoryImpl{}
}

func (self *CafeRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, cafe CafeDatabaseIO) CafeDatabaseIO {
	// Insert new cafe
	SQLPut := "INSERT INTO cafe(email, title, password) VALUES ($1,$2,$3) RETURNING id;"
	tx.QueryRowContext(ctx, SQLPut, cafe.Email, cafe.Title, cafe.Password).Scan(&cafe.Id)

	// Return created cafe
	return cafe
}

func (self *CafeRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, cafe CafeDatabaseIO) CafeDatabaseIO {
	// Update existing cafe
	SQLPut := "UPDATE cafe SET email=$1, title=$2, password=$3, updated_at=NOW() WHERE id=$4 AND is_active=true;"
	_, err := tx.ExecContext(ctx, SQLPut, cafe.Email, cafe.Title, cafe.Password, cafe.Id)
	helper.PanicIfError(err)

	// Return updated cafe
	cafe, err = self.FindById(ctx, tx, cafe.Id)
	helper.PanicIfError(err)
	return cafe
}

func (self *CafeRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) {
	// Delete existing cafe
	SQLDel := "UPDATE cafe SET deleted_at=NOW(), is_active=TRUE WHERE id=$1;"
	_, err := tx.ExecContext(ctx, SQLDel, id)
	helper.PanicIfError(err)
}

func (self *CafeRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (CafeDatabaseIO, error) {
	// Extract existing cafe
	SQLGet := "SELECT id, email, title, password, created_at, updated_at FROM cafe WHERE email=$1 AND is_active=true;"
	rows, err := tx.QueryContext(ctx, SQLGet, email)
	helper.PanicIfError(err)

	// Bind all columns value to cafe variable
	cafe := CafeDatabaseIO{}
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&cafe.Id, &cafe.Email, &cafe.Title, &cafe.Password, &cafe.CreatedAt, &cafe.UpdatedAt)
		helper.PanicIfError(err)
		return cafe, nil
	} else {
		return cafe, errors.New("Cafe not found")
	}
}

func (self *CafeRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (CafeDatabaseIO, error) {
	// Extract existing cafe
	SQLGet := "SELECT email, title, password, created_at, updated_at FROM cafe WHERE id=$1 AND is_active=true LIMIT 1;"
	rows, err := tx.QueryContext(ctx, SQLGet, id)
	helper.PanicIfError(err)

	// Bind all columns value to cafe variable
	cafe := CafeDatabaseIO{}
	cafe.Id = id
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&cafe.Email, &cafe.Title, &cafe.Password, &cafe.CreatedAt, &cafe.UpdatedAt)
		helper.PanicIfError(err)
		return cafe, nil
	} else {
		return cafe, errors.New("Cafe not found")
	}
}

func (self *CafeRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []CafeDatabaseIO {
	// Extract all cafe
	SQLGet := "SELECT id, email, title, password, created_at, updated_at, is_active FROM cafe WHERE is_active=true"
	rows, err := tx.QueryContext(ctx, SQLGet)
	helper.PanicIfError(err)

	// Iterate all extracted rows
	var listCafe []CafeDatabaseIO
	defer rows.Close()
	for rows.Next() {
		cafe := CafeDatabaseIO{}
		err := rows.Scan(&cafe.Id, &cafe.Email, &cafe.Title, &cafe.Password, &cafe.CreatedAt, &cafe.UpdatedAt, &cafe.IsActive)
		helper.PanicIfError(err)

		listCafe = append(listCafe, cafe)
	}

	return listCafe
}
