package user

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/michaelact/app"
)

func TestUserRepositoryImpl_Insert(t *testing.T) {
	db = app.NewDB()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("Failed to begin transaction: %v", err)
	}
	defer tx.Rollback()

	userRepo := NewUserRepository()
	user := UserDatabaseIO{
		Email:    "test@example.com",
		Password: "password123",
	}

	// Call the Insert method
	insertedUser := userRepo.Insert(ctx, tx, user)

	// Add assertions to check if the user was inserted correctly
	if insertedUser.Id == 0 {
		t.Error("Insert did not return a valid ID for the user")
	}

	// You can also check other properties of the inserted user

	// Commit the transaction if everything is successful
	tx.Commit()
}

func TestUserRepositoryImpl_Update(t *testing.T) {
	// Set up your database connection and transaction here (you can use a testing database or a mock)
	// Initialize your UserRepository implementation

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("Failed to begin transaction: %v", err)
	}
	defer tx.Rollback()

	userRepo := NewUserRepository()
	user := UserDatabaseIO{
		Id:       1, // Assuming this ID exists in your test database
		Email:    "updated@example.com",
		Password: "newpassword",
	}

	// Call the Update method
	updatedUser := userRepo.Update(ctx, tx, user)

	// Add assertions to check if the user was updated correctly

	// Commit the transaction if everything is successful
	tx.Commit()
}

func TestUserRepositoryImpl_Delete(t *testing.T) {
	// Set up your database connection and transaction here (you can use a testing database or a mock)
	// Initialize your UserRepository implementation

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("Failed to begin transaction: %v", err)
	}
	defer tx.Rollback()

	userRepo := NewUserRepository()
	userID := 1 // Assuming this ID exists in your test database

	// Call the Delete method
	userRepo.Delete(ctx, tx, userID)

	// Add assertions to check if the user was marked as deleted correctly

	// Commit the transaction if everything is successful
	tx.Commit()
}

func TestUserRepositoryImpl_FindById(t *testing.T) {
	// Set up your database connection and transaction here (you can use a testing database or a mock)
	// Initialize your UserRepository implementation

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("Failed to begin transaction: %v", err)
	}
	defer tx.Rollback()

	userRepo := NewUserRepository()
	userID := 1 // Assuming this ID exists in your test database

	// Call the FindById method
	foundUser, err := userRepo.FindById(ctx, tx, userID)

	// Add assertions to check if the user was found correctly and check its properties

	if err != nil {
		t.Fatalf("Error finding user by ID: %v", err)
	}

	// Commit the transaction if everything is successful
	tx.Commit()
}

func TestUserRepositoryImpl_FindAll(t *testing.T) {
	// Set up your database connection and transaction here (you can use a testing database or a mock)
	// Initialize your UserRepository implementation

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("Failed to begin transaction: %v", err)
	}
	defer tx.Rollback()

	userRepo := NewUserRepository()

	// Call the FindAll method
	users := userRepo.FindAll(ctx, tx, 1) // Assuming '1' is the page you want to fetch

	// Add assertions to check if users were retrieved correctly

	// Commit the transaction if everything is successful
	tx.Commit()
}
