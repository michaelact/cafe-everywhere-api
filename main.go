package main

import (
	"github.com/go-playground/validator/v10"
	"net/http"

	"github.com/michaelact/cafe-everywhere/module/user"
	"github.com/michaelact/cafe-everywhere/middleware"
	"github.com/michaelact/cafe-everywhere/helper"
	"github.com/michaelact/cafe-everywhere/app"
	
)

func InitializeServer() *http.Server {
	conf := app.NewConfig()
	db := app.NewDB(conf)
	validate := validator.New()

	userRepository := user.NewUserRepository()
	userService := user.NewUserService(userRepository, db, validate)
	userController := user.NewUserController(userService)

	router := app.NewRouter(userController)
	authMiddleware := middleware.NewAuthMiddleware(router, conf)

	server := app.NewServer(authMiddleware, conf)
	return server
}

func main() {
	server := InitializeServer()
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
