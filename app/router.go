package app

import (
	"github.com/julienschmidt/httprouter"

	"github.com/michaelact/cafe-everywhere/module/order"
	"github.com/michaelact/cafe-everywhere/module/menu"
	"github.com/michaelact/cafe-everywhere/module/cafe"
	"github.com/michaelact/cafe-everywhere/module/user"
	"github.com/michaelact/cafe-everywhere/exception"
)

func NewRouter(userController user.UserController, cafeController cafe.CafeController, menuController menu.MenuController, orderController order.OrderController) *httprouter.Router {
	router := httprouter.New()
	router.PanicHandler = exception.ErrorHandler

	router.POST("/users/login", userController.Login)
	router.GET("/users", userController.FindAll)
	router.POST("/users", userController.Create)
	router.GET("/users/:userId", userController.FindById)
	router.PATCH("/users/:userId", userController.Update)
	router.DELETE("/users/:userId", userController.Delete)
	router.GET("/users-order/:userId", orderController.FindByUserId)

	router.POST("/cafe/login", cafeController.Login)
	router.GET("/cafe", cafeController.FindAll)
	router.POST("/cafe", cafeController.Create)
	router.GET("/cafe/:cafeId", cafeController.FindById)
	router.PATCH("/cafe/:cafeId", cafeController.Update)
	router.DELETE("/cafe/:cafeId", cafeController.Delete)
	router.GET("/cafe-menu/:cafeId", menuController.FindByCafeId)

	router.GET("/menu", menuController.FindAll)
	router.POST("/menu", menuController.Create)
	router.GET("/menu/:menuId", menuController.FindById)
	router.PATCH("/menu/:menuId", menuController.Update)
	router.DELETE("/menu/:menuId", menuController.Delete)

	router.GET("/order", orderController.FindAll)
	router.POST("/order", orderController.Create)
	router.GET("/order/:orderId", orderController.FindById)
	router.PATCH("/order/:orderId", orderController.Update)
	router.DELETE("/order/:orderId", orderController.Delete)

	return router
}
