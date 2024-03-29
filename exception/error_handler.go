package exception

import (
	"net/http"
	"log"

	"github.com/go-playground/validator/v10"

	"github.com/michaelact/cafe-everywhere/helper"
)

func ErrorHandler(res http.ResponseWriter, req *http.Request, err interface{}) {
	if notFoundError(res, req, err) { return }
	if validationError(res, req, err) { return }

	// If error doesn't match the error above
	log.Println(err)
	message := "Something went wrong!"
	helper.WriteToResponseBodyError(res, http.StatusInternalServerError, message)
}

func validationError(res http.ResponseWriter, req *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		helper.WriteToResponseBodyError(res, http.StatusBadRequest, exception.Error())
		return true
	} else {
		return false
	}
}

func notFoundError(res http.ResponseWriter, req *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		helper.WriteToResponseBodyError(res, http.StatusNotFound, exception.Error())
		return true
	} else {
		return false
	}
}
