package menu

import (
	"strconv"
	"net/http"
	"github.com/julienschmidt/httprouter"

	"github.com/michaelact/cafe-everywhere/helper"
)

type MenuController interface {
	Create(res http.ResponseWriter, req *http.Request, params httprouter.Params)
	Update(res http.ResponseWriter, req *http.Request, params httprouter.Params)
	Delete(res http.ResponseWriter, req *http.Request, params httprouter.Params)
	FindById(res http.ResponseWriter, req *http.Request, params httprouter.Params)
	FindByCafeId(res http.ResponseWriter, req *http.Request, params httprouter.Params)
	FindAll(res http.ResponseWriter, req *http.Request, params httprouter.Params)
}

type MenuControllerImpl struct {
	MenuService MenuService
}

func NewMenuController(menuService MenuService) MenuController {
	return &MenuControllerImpl{
		MenuService: menuService, 
	}
}

func (self *MenuControllerImpl) Create(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	request := HTTPMenuRequest{}
	helper.ReadFromRequestBody(req, &request)

	serviceResponse := self.MenuService.Create(req.Context(), request)
	webResponse := helper.WebResponse{
		Status:  "Success", 
		Message: "Success", 
		Data:    serviceResponse, 
	}

	helper.WriteToResponseBody(res, &webResponse)
}

func (self *MenuControllerImpl) Update(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	request := HTTPMenuUpdateRequest{}
	helper.ReadFromRequestBody(req, &request)

	// Bind ID from Route Parameter
	id, err := strconv.Atoi(params.ByName("menuId"))
	helper.PanicIfError(err)
	request.Id = id

	serviceResponse := self.MenuService.Update(req.Context(), request)
	webResponse := helper.WebResponse{
		Status:  "Success", 
		Message: "Success", 
		Data:    serviceResponse, 
	}

	helper.WriteToResponseBody(res, &webResponse)
}

func (self *MenuControllerImpl) Delete(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	requestId, err := strconv.Atoi(params.ByName("menuId"))
	helper.PanicIfError(err)

	self.MenuService.Delete(req.Context(), requestId)
	webResponse := helper.WebResponse{
		Status:  "Success", 
		Message: "Success", 
	}

	helper.WriteToResponseBody(res, &webResponse)
}

func (self *MenuControllerImpl) FindByCafeId(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	requestId, err := strconv.Atoi(params.ByName("cafeId"))
	helper.PanicIfError(err)

	serviceResponse := self.MenuService.FindByCafeId(req.Context(), requestId)
	webResponse := helper.WebResponse{
		Status:  "Success", 
		Message: "Success", 
		Data:    serviceResponse, 
	}

	helper.WriteToResponseBody(res, &webResponse)
}

func (self *MenuControllerImpl) FindById(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	requestId, err := strconv.Atoi(params.ByName("menuId"))
	helper.PanicIfError(err)

	serviceResponse := self.MenuService.FindById(req.Context(), requestId)
	webResponse := helper.WebResponse{
		Status:  "Success", 
		Message: "Success", 
		Data:    serviceResponse, 
	}

	helper.WriteToResponseBody(res, &webResponse)
}

func (self *MenuControllerImpl) FindAll(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	serviceResponse := self.MenuService.FindAll(req.Context())
	webResponse := helper.WebResponse{
		Status:  "Success", 
		Message: "Success", 
		Data:    serviceResponse, 
	}

	helper.WriteToResponseBody(res, &webResponse)
}
