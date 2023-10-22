package cafe

import (
	"strconv"
	"net/http"
	"github.com/julienschmidt/httprouter"

	"github.com/michaelact/cafe-everywhere/helper"
)

type CafeController interface {
	Login(res http.ResponseWriter, req *http.Request, params httprouter.Params)
	Create(res http.ResponseWriter, req *http.Request, params httprouter.Params)
	Update(res http.ResponseWriter, req *http.Request, params httprouter.Params)
	Delete(res http.ResponseWriter, req *http.Request, params httprouter.Params)
	FindById(res http.ResponseWriter, req *http.Request, params httprouter.Params)
	FindAll(res http.ResponseWriter, req *http.Request, params httprouter.Params)
}

type CafeControllerImpl struct {
	CafeService CafeService
}

func NewCafeController(cafeService CafeService) CafeController {
	return &CafeControllerImpl{
		CafeService: cafeService, 
	}
}

func (self *CafeControllerImpl) Login(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	request := HTTPCafeRequest{}
	helper.ReadFromRequestBody(req, &request)

	var webResponse helper.WebResponse
	serviceResponse, ok := self.CafeService.Login(req.Context(), request)
	if ok {
		webResponse = helper.WebResponse{
			Status:  "Success", 
			Message: "Cafe Sign-in Success.", 
			Data:    serviceResponse,
		}
	} else {
		webResponse = helper.WebResponse{
			Status:  "Failed", 
			Message: "Cafe Sign-in Failure.", 
		}
	}

	helper.WriteToResponseBody(res, &webResponse)
}

func (self *CafeControllerImpl) Create(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	request := HTTPCafeRequest{}
	helper.ReadFromRequestBody(req, &request)

	serviceResponse := self.CafeService.Create(req.Context(), request)
	webResponse := helper.WebResponse{
		Status:  "Success", 
		Message: "Success", 
		Data:    serviceResponse, 
	}

	helper.WriteToResponseBody(res, &webResponse)
}

func (self *CafeControllerImpl) Update(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	request := HTTPCafeUpdateRequest{}
	helper.ReadFromRequestBody(req, &request)

	// Bind ID from Route Parameter
	id, err := strconv.Atoi(params.ByName("cafeId"))
	helper.PanicIfError(err)
	request.Id = id

	serviceResponse := self.CafeService.Update(req.Context(), request)
	webResponse := helper.WebResponse{
		Status:  "Success", 
		Message: "Success", 
		Data:    serviceResponse, 
	}

	helper.WriteToResponseBody(res, &webResponse)
}

func (self *CafeControllerImpl) Delete(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	requestId, err := strconv.Atoi(params.ByName("cafeId"))
	helper.PanicIfError(err)

	self.CafeService.Delete(req.Context(), requestId)
	webResponse := helper.WebResponse{
		Status:  "Success", 
		Message: "Success", 
	}

	helper.WriteToResponseBody(res, &webResponse)
}

func (self *CafeControllerImpl) FindById(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	requestId, err := strconv.Atoi(params.ByName("cafeId"))
	helper.PanicIfError(err)

	serviceResponse := self.CafeService.FindById(req.Context(), requestId)
	webResponse := helper.WebResponse{
		Status:  "Success", 
		Message: "Success", 
		Data:    serviceResponse, 
	}

	helper.WriteToResponseBody(res, &webResponse)
}

func (self *CafeControllerImpl) FindAll(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	serviceResponse := self.CafeService.FindAll(req.Context())
	webResponse := helper.WebResponse{
		Status:  "Success", 
		Message: "Success", 
		Data:    serviceResponse, 
	}

	helper.WriteToResponseBody(res, &webResponse)
}