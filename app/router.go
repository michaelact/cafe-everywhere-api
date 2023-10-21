package app

import (
	"github.com/julienschmidt/httprouter"

	"github.com/michaelact/cafe-everywhere/module/user"
	"github.com/michaelact/cafe-everywhere/exception"
)

func NewRouter(userController user.UserController) *httprouter.Router {
	router := httprouter.New()
	router.PanicHandler = exception.ErrorHandler

	router.POST("/users/login", userController.Login)
	router.GET("/users", userController.FindAll)
	router.POST("/users", userController.Create)
	router.GET("/users/:userId", userController.FindById)
	router.PATCH("/users/:userId", userController.Update)
	router.DELETE("/users/:userId", userController.Delete)	

	return router
}
