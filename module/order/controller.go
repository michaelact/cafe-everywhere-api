package order

import (
	"strconv"
	"net/http"
	"github.com/julienschmidt/httprouter"

	"github.com/michaelact/cafe-everywhere/helper"
)

type OrderController interface {
	Create(res http.ResponseWriter, req *http.Request, params httprouter.Params)
	Update(res http.ResponseWriter, req *http.Request, params httprouter.Params)
	Delete(res http.ResponseWriter, req *http.Request, params httprouter.Params)
	FindById(res http.ResponseWriter, req *http.Request, params httprouter.Params)
	FindByCafeId(res http.ResponseWriter, req *http.Request, params httprouter.Params)
	FindByUserId(res http.ResponseWriter, req *http.Request, params httprouter.Params)
	FindAll(res http.ResponseWriter, req *http.Request, params httprouter.Params)
}

type OrderControllerImpl struct {
	OrderService OrderService
}

func NewOrderController(orderService OrderService) OrderController {
	return &OrderControllerImpl{
		OrderService: orderService, 
	}
}

func (self *OrderControllerImpl) Create(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	request := HTTPOrderRequest{}
	helper.ReadFromRequestBody(req, &request)

	serviceResponse := self.OrderService.Create(req.Context(), request)
	webResponse := helper.WebResponse{
		Status:  "Success", 
		Message: "Success", 
		Data:    serviceResponse, 
	}

	helper.WriteToResponseBody(res, &webResponse)
}

func (self *OrderControllerImpl) Update(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	request := HTTPOrderUpdateRequest{}
	helper.ReadFromRequestBody(req, &request)

	// Bind ID from Route Parameter
	id, err := strconv.Atoi(params.ByName("orderId"))
	helper.PanicIfError(err)
	request.Id = id

	serviceResponse := self.OrderService.Update(req.Context(), request)
	webResponse := helper.WebResponse{
		Status:  "Success", 
		Message: "Success", 
		Data:    serviceResponse, 
	}

	helper.WriteToResponseBody(res, &webResponse)
}

func (self *OrderControllerImpl) Delete(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	requestId, err := strconv.Atoi(params.ByName("orderId"))
	helper.PanicIfError(err)

	self.OrderService.Delete(req.Context(), requestId)
	webResponse := helper.WebResponse{
		Status:  "Success", 
		Message: "Success", 
	}

	helper.WriteToResponseBody(res, &webResponse)
}

func (self *OrderControllerImpl) FindByCafeId(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	requestId, err := strconv.Atoi(params.ByName("cafeId"))
	helper.PanicIfError(err)

	serviceResponse := self.OrderService.FindByCafeId(req.Context(), requestId)
	webResponse := helper.WebResponse{
		Status:  "Success", 
		Message: "Success", 
		Data:    serviceResponse, 
	}

	helper.WriteToResponseBody(res, &webResponse)
}

func (self *OrderControllerImpl) FindByUserId(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	requestId, err := strconv.Atoi(params.ByName("userId"))
	helper.PanicIfError(err)

	serviceResponse := self.OrderService.FindByUserId(req.Context(), requestId)
	webResponse := helper.WebResponse{
		Status:  "Success", 
		Message: "Success", 
		Data:    serviceResponse, 
	}

	helper.WriteToResponseBody(res, &webResponse)
}

func (self *OrderControllerImpl) FindById(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	requestId, err := strconv.Atoi(params.ByName("orderId"))
	helper.PanicIfError(err)

	serviceResponse := self.OrderService.FindById(req.Context(), requestId)
	webResponse := helper.WebResponse{
		Status:  "Success", 
		Message: "Success", 
		Data:    serviceResponse, 
	}

	helper.WriteToResponseBody(res, &webResponse)
}

func (self *OrderControllerImpl) FindAll(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	serviceResponse := self.OrderService.FindAll(req.Context())
	webResponse := helper.WebResponse{
		Status:  "Success", 
		Message: "Success", 
		Data:    serviceResponse, 
	}

	helper.WriteToResponseBody(res, &webResponse)
}
