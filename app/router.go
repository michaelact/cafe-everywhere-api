package app

import (
	"github.com/julienschmidt/httprouter"

	"github.com/michaelact/cafe-everywhere/module/cafe"
	"github.com/michaelact/cafe-everywhere/module/user"
	"github.com/michaelact/cafe-everywhere/exception"
)

func NewRouter(userController user.UserController, cafeController cafe.CafeController) *httprouter.Router {
	router := httprouter.New()
	router.PanicHandler = exception.ErrorHandler

	router.POST("/users/login", userController.Login)
	router.GET("/users", userController.FindAll)
	router.POST("/users", userController.Create)
	router.GET("/users/:userId", userController.FindById)
	router.PATCH("/users/:userId", userController.Update)
	router.DELETE("/users/:userId", userController.Delete)

	router.POST("/cafe/login", cafeController.Login)
	router.GET("/cafe", cafeController.FindAll)
	router.POST("/cafe", cafeController.Create)
	router.GET("/cafe/:cafeId", cafeController.FindById)
	router.PATCH("/cafe/:cafeId", cafeController.Update)
	router.DELETE("/cafe/:cafeId", cafeController.Delete)

	return router
}
