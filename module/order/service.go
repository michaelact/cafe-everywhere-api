package order

import(
	"github.com/go-playground/validator/v10"
	"database/sql"
	"context"
	"log"

	"github.com/michaelact/cafe-everywhere/exception"
	"github.com/michaelact/cafe-everywhere/helper"
)

type OrderService interface {
	Create(ctx context.Context, request HTTPOrderRequest) HTTPOrderResponse
	Update(ctx context.Context, request HTTPOrderUpdateRequest) HTTPOrderResponse
	Delete(ctx context.Context, id int)
	FindById(ctx context.Context, id int) HTTPOrderResponse
	FindByCafeId(ctx context.Context, cafeId int) []HTTPOrderResponse
	FindByUserId(ctx context.Context, userId int) []HTTPOrderResponse
	FindAll(ctx context.Context) []HTTPOrderResponse
}

type OrderServiceImpl struct {
	OrderRepository OrderRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewOrderService(OrderRepository OrderRepository, db *sql.DB, validate *validator.Validate) OrderService {
	return &OrderServiceImpl{
		OrderRepository: OrderRepository, 
		DB:             db, 
		Validate:       validate, 
	}
}

func (self *OrderServiceImpl) Create(ctx context.Context, request HTTPOrderRequest) HTTPOrderResponse {
	// Fail if request is not valid
	err := self.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := self.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	order := OrderDatabaseIO{
		MenuId:  request.MenuId,
		UserId:  request.UserId,
		Count:   request.Count,
		Notes:   request.Notes,
		Address: request.Address,
	}

	order = self.OrderRepository.Insert(ctx, tx, order)
	return ToOrderResponse(order)
}

func (self *OrderServiceImpl) Update(ctx context.Context, request HTTPOrderUpdateRequest) HTTPOrderResponse {
	// Fail if request is not valid
	err := self.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := self.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Fail if existing order not found
	order, err := self.OrderRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	// Update existing order
	order.Status = request.Status
	order = self.OrderRepository.Update(ctx, tx, order)
	return ToOrderResponse(order)
}

func (self *OrderServiceImpl) Delete(ctx context.Context, id int) {
	tx, err := self.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Fail if existing order not found
	_, err = self.OrderRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	// Delete existing order
	self.OrderRepository.Delete(ctx, tx, id)
}

func (self *OrderServiceImpl) FindByCafeId(ctx context.Context, userId int) []HTTPOrderResponse {
	tx, err := self.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	listOrder := self.OrderRepository.FindByCafeId(ctx, tx, userId)
	return ToOrderResponses(listOrder)
}

func (self *OrderServiceImpl) FindByUserId(ctx context.Context, userId int) []HTTPOrderResponse {
	tx, err := self.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	listOrder := self.OrderRepository.FindByUserId(ctx, tx, userId)
	return ToOrderResponses(listOrder)
}

func (self *OrderServiceImpl) FindById(ctx context.Context, id int) HTTPOrderResponse {
	tx, err := self.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Fail if existing order not found
	order, err := self.OrderRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return ToOrderResponse(order)
}

func (self *OrderServiceImpl) FindAll(ctx context.Context) []HTTPOrderResponse {
	tx, err := self.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	listOrder := self.OrderRepository.FindAll(ctx, tx)
	log.Println(listOrder)
	return ToOrderResponses(listOrder)
}
