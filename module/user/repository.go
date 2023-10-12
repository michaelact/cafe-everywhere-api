package user

import (
	"database/sql"
	"context"
	"errors"

	"github.com/michaelact/cafe-everywhere/helper"
)

type UserRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, user UserDatabaseIO) UserDatabaseIO
	Update(ctx context.Context, tx *sql.Tx, user UserDatabaseIO) UserDatabaseIO
	Delete(ctx context.Context, tx *sql.Tx, id int)
	FindById(ctx context.Context, tx *sql.Tx, id int) (UserDatabaseIO, error)
	FindAll(ctx context.Context, tx *sql.Tx, page int) []UserDatabaseIO
}

type UserRepositoryImpl struct {}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (self *UserRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, user UserDatabaseIO) UserDatabaseIO {
	// Insert new user
	SQLPut := "INSERT INTO users(email, password) VALUES ($1,$2) RETURNING id;"
	tx.QueryRowContext(ctx, SQLPut, user.Email, user.Password).Scan(&user.Id)

	// Return created user
	return user
}

func (self *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user UserDatabaseIO) UserDatabaseIO {
	// Update existing user
	SQLPut := "UPDATE users SET email=$1, password=$2 WHERE id=$3;"
	_, err := tx.ExecContext(ctx, SQLPut, user.Email, user.Password, user.Id)
	helper.PanicIfError(err)

	// Return updated user
	user, err = self.FindById(ctx, tx, user.Id)
	helper.PanicIfError(err)
	return user
}

func (self *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) {
	// Delete existing user
	SQLDel := "UPDATE users SET deleted_at=NOW(), is_active=TRUE WHERE id=$1;"
	_, err := tx.ExecContext(ctx, SQLDel, id)
	helper.PanicIfError(err)
}

func (self *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (UserDatabaseIO, error) {
	// Extract existing user
	SQLGet := "SELECT email, password, created_at, updated_at, deleted_at, is_active FROM users WHERE id=$1;"
	rows, err := tx.QueryContext(ctx, SQLGet, id)
	helper.PanicIfError(err)

	// Bind all columns value to user variable
	user := UserDatabaseIO{}
	user.Id = id
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt, &user.IsActive)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("User not found")
	}
}

func (self *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, page int) []UserDatabaseIO {
	// Extract all user
	SQLGet := "SELECT email, password, created_at, updated_at, deleted_at, is_active FROM users"
	rows, err := tx.QueryContext(ctx, SQLGet)
	helper.PanicIfError(err)

	// Iterate all extracted rows
	var listUser []UserDatabaseIO
	defer rows.Close()
	for rows.Next() {
		user := UserDatabaseIO{}
		err := rows.Scan(&user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt, &user.IsActive)
		helper.PanicIfError(err)

		listUser = append(listUser, user)
	}

	return listUser
}
