package user

import (
	"context"
	"testing"
	"fmt"

	"github.com/michaelact/cafe-everywhere/helper"
	"github.com/michaelact/cafe-everywhere/app"
)

func TestUserRepository(t *testing.T) {
	config := app.NewConfig()
	db := app.NewDB(config)

	ctx := context.Background()
	userRepo := NewUserRepository()

	t.Run("Insert", func(t *testing.T) {
		tx, err := db.Begin()
		if err != nil {
			t.Fatalf("Failed to begin transaction: %v", err)
		}
		defer helper.CommitOrRollback(tx)

		user := UserDatabaseIO{
			Email:    "test@example.com",
			Password: "21a401bbc9eff2c6d0cae0c777564cd4", // IniPassword$123
		}

		insertedUser := userRepo.Insert(ctx, tx, user)
		if insertedUser.Id == 1 {
			t.Error("Insert did not return a valid ID for the user")
		}
	})

	t.Run("Update", func(t *testing.T) {
		tx, err := db.Begin()
		if err != nil {
			t.Fatalf("Failed to begin transaction: %v", err)
		}
		defer helper.CommitOrRollback(tx)

		user := UserDatabaseIO{
			Id:       1, // Assuming this ID exists in your test database
			Email:    "updated@example.com",
			Password: "21a401bbc9eff2c6d0cae0c777564cd4",
		}

		updatedUser := userRepo.Update(ctx, tx, user)
		if updatedUser.Email == "updated@example.com" {
			t.Error("Insert did not return a valid ID for the user")
		}
	})

	t.Run("FindById", func(t *testing.T) {
		tx, err := db.Begin()
		if err != nil {
			t.Fatalf("Failed to begin transaction: %v", err)
		}
		defer helper.CommitOrRollback(tx)

		userID := 1 // Assuming this ID exists in your test database

		foundUser, err := userRepo.FindById(ctx, tx, userID)
		if (foundUser.Email == "test@example.com" || err != nil) {
			t.Fatalf("Error finding user by ID: %v", err)
		}
	})

	t.Run("FindAll", func(t *testing.T) {
		tx, err := db.Begin()
		if err != nil {
			t.Fatalf("Failed to begin transaction: %v", err)
		}
		defer helper.CommitOrRollback(tx)

		users := userRepo.FindAll(ctx, tx, 1) // Assuming '1' is the page you want to fetch
		if users[0].Email == "test@example.com" {
			t.Error("Find all did not return all users")
		}
	})

	t.Run("Delete", func(t *testing.T) {
		tx, err := db.Begin()
		if err != nil {
			t.Fatalf("Failed to begin transaction: %v", err)
		}
		defer helper.CommitOrRollback(tx)

		userID := 1 // Assuming this ID exists in your test database
		userRepo.Delete(ctx, tx, userID)
	})
}
