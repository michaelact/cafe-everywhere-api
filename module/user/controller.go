package user

import (
	"strconv"
	"net/http"
	"github.com/julienschmidt/httprouter"

	"github.com/michaelact/cafe-everywhere/helper"
)

type UserController interface {
	Login(res http.ResponseWriter, req *http.Request, params httprouter.Params)
	Create(res http.ResponseWriter, req *http.Request, params httprouter.Params)
	Update(res http.ResponseWriter, req *http.Request, params httprouter.Params)
	Delete(res http.ResponseWriter, req *http.Request, params httprouter.Params)
	FindById(res http.ResponseWriter, req *http.Request, params httprouter.Params)
	FindAll(res http.ResponseWriter, req *http.Request, params httprouter.Params)
}

type UserControllerImpl struct {
	UserService UserService
}

func NewUserController(userService UserService) UserController {
	return &UserControllerImpl{
		UserService: userService, 
	}
}

func (self *UserControllerImpl) Login(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	request := HTTPUserRequest{}
	helper.ReadFromRequestBody(req, &request)

	var webResponse helper.WebResponse
	serviceResponse, ok := self.UserService.Login(req.Context(), request)
	if ok {
		webResponse = helper.WebResponse{
			Status:  "Success", 
			Message: "User Sign-in Success.", 
			Data:    serviceResponse,
		}

		helper.WriteToResponseBody(res, &webResponse)
	} else {
		helper.WriteToResponseBodyError(res, http.StatusUnauthorized, "User Sign-in Failure.")
	}
}

func (self *UserControllerImpl) Create(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	request := HTTPUserRequest{}
	helper.ReadFromRequestBody(req, &request)

	serviceResponse := self.UserService.Create(req.Context(), request)
	webResponse := helper.WebResponse{
		Status:  "Success", 
		Message: "Success", 
		Data:    serviceResponse, 
	}

	helper.WriteToResponseBody(res, &webResponse)
}

func (self *UserControllerImpl) Update(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	request := HTTPUserUpdateRequest{}
	helper.ReadFromRequestBody(req, &request)

	// Bind ID from Route Parameter
	id, err := strconv.Atoi(params.ByName("userId"))
	helper.PanicIfError(err)
	request.Id = id

	serviceResponse := self.UserService.Update(req.Context(), request)
	webResponse := helper.WebResponse{
		Status:  "Success", 
		Message: "Success", 
		Data:    serviceResponse, 
	}

	helper.WriteToResponseBody(res, &webResponse)
}

func (self *UserControllerImpl) Delete(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	requestId, err := strconv.Atoi(params.ByName("userId"))
	helper.PanicIfError(err)

	self.UserService.Delete(req.Context(), requestId)
	webResponse := helper.WebResponse{
		Status:  "Success", 
		Message: "Success", 
	}

	helper.WriteToResponseBody(res, &webResponse)
}

func (self *UserControllerImpl) FindById(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	requestId, err := strconv.Atoi(params.ByName("userId"))
	helper.PanicIfError(err)

	serviceResponse := self.UserService.FindById(req.Context(), requestId)
	webResponse := helper.WebResponse{
		Status:  "Success", 
		Message: "Success", 
		Data:    serviceResponse, 
	}

	helper.WriteToResponseBody(res, &webResponse)
}

func (self *UserControllerImpl) FindAll(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	serviceResponse := self.UserService.FindAll(req.Context())
	webResponse := helper.WebResponse{
		Status:  "Success", 
		Message: "Success", 
		Data:    serviceResponse, 
	}

	helper.WriteToResponseBody(res, &webResponse)
}