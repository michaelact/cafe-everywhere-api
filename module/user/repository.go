package user

import (
	"database/sql"
	"context"
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
	SQLPut := "INSERT INTO Users(email, password) VALUES(email, password)"
	result, err := tx.ExecContext(ctx, SQLPut, user.Email, user.Password)
	helper.PanicIfError(err)

	// Return created user
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	user, err = self.FindById(ctx, tx, int(id))
	helper.PanicIfError(err)
	return user
}

func (self *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user UserDatabaseIO) UserDatabaseIO {
	// Update existing user
	SQLPut := "UPDATE Users SET email=?, password=? WHERE id=?"
	_, err := tx.ExecContext(ctx, SQLPut, user.Title, user.Email, user.Id)
	helper.PanicIfError(err)

	// Return updated user
	user, err = self.FindById(ctx, tx, user.Id)
	helper.PanicIfError(err)
	return user
}

func (self *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) {
	// Delete existing user
	SQLDel := "UPDATE Users SET deleted_at=NOW(), is_active=TRUE WHERE id=?"
	_, err := tx.ExecContext(ctx, SQLDel, id)
	helper.PanicIfError(err)
}

func (self *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) UserDatabaseIO {
	// Extract existing user
	SQLGet := "SELECT email, password, created_at, updated_at, deleted_at, is_active FROM Users WHERE id=?"
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
	SQLGet := "SELECT email, password, created_at, updated_at, deleted_at, is_active FROM Users"
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
