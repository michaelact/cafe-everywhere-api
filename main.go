package main

import (
	"github.com/go-playground/validator/v10"
	"net/http"

	"github.com/michaelact/cafe-everywhere/module/order"
	"github.com/michaelact/cafe-everywhere/module/menu"
	"github.com/michaelact/cafe-everywhere/module/cafe"
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

	cafeRepository := cafe.NewCafeRepository()
	cafeService := cafe.NewCafeService(cafeRepository, db, validate)
	cafeController := cafe.NewCafeController(cafeService)

	menuRepository := menu.NewMenuRepository()
	menuService := menu.NewMenuService(menuRepository, db, validate)
	menuController := menu.NewMenuController(menuService)

	orderRepository := order.NewOrderRepository()
	orderService := order.NewOrderService(orderRepository, db, validate)
	orderController := order.NewOrderController(orderService)

	router := app.NewRouter(userController, cafeController, menuController, orderController)
	authMiddleware := middleware.NewAuthMiddleware(router, conf)

	server := app.NewServer(authMiddleware, conf)
	return server
}

func main() {
	server := InitializeServer()
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
