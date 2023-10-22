package menu

import(
	"github.com/go-playground/validator/v10"
	"database/sql"
	"context"
	"log"

	"github.com/michaelact/cafe-everywhere/exception"
	"github.com/michaelact/cafe-everywhere/helper"
)

type MenuService interface {
	Create(ctx context.Context, request HTTPMenuRequest) HTTPMenuResponse
	Update(ctx context.Context, request HTTPMenuUpdateRequest) HTTPMenuResponse
	Delete(ctx context.Context, id int)
	FindById(ctx context.Context, id int) HTTPMenuResponse
	FindByCafeId(ctx context.Context, cafeId int) []HTTPMenuResponse
	FindAll(ctx context.Context) []HTTPMenuResponse
}

type MenuServiceImpl struct {
	MenuRepository MenuRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewMenuService(MenuRepository MenuRepository, db *sql.DB, validate *validator.Validate) MenuService {
	return &MenuServiceImpl{
		MenuRepository: MenuRepository, 
		DB:             db, 
		Validate:       validate, 
	}
}

func (self *MenuServiceImpl) Create(ctx context.Context, request HTTPMenuRequest) HTTPMenuResponse {
	// Fail if request is not valid
	err := self.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := self.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	menu := MenuDatabaseIO{
		Title:       request.Title,
		Description: request.Description,
		Count:       request.Count,
		Price:       request.Price,
		CafeId:      request.CafeId,
	}

	menu = self.MenuRepository.Insert(ctx, tx, menu)
	return ToMenuResponse(menu)
}

func (self *MenuServiceImpl) Update(ctx context.Context, request HTTPMenuUpdateRequest) HTTPMenuResponse {
	// Fail if request is not valid
	err := self.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := self.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Fail if existing menu not found
	menu, err := self.MenuRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	// Update existing menu
	menu.Title = request.Title
	menu.Description = request.Description
	menu.Count = request.Count
	menu.Price = request.Price
	menu.CafeId = request.CafeId
	menu = self.MenuRepository.Update(ctx, tx, menu)
	return ToMenuResponse(menu)
}

func (self *MenuServiceImpl) Delete(ctx context.Context, id int) {
	tx, err := self.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Fail if existing menu not found
	_, err = self.MenuRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	// Delete existing menu
	self.MenuRepository.Delete(ctx, tx, id)
}

func (self *MenuServiceImpl) FindByCafeId(ctx context.Context, cafeId int) []HTTPMenuResponse {
	tx, err := self.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Fail if existing menu not found
	listMenu := self.MenuRepository.FindByCafeId(ctx, tx, cafeId)
	return ToMenuResponses(listMenu)
}

func (self *MenuServiceImpl) FindById(ctx context.Context, id int) HTTPMenuResponse {
	tx, err := self.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Fail if existing menu not found
	menu, err := self.MenuRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return ToMenuResponse(menu)
}

func (self *MenuServiceImpl) FindAll(ctx context.Context) []HTTPMenuResponse {
	tx, err := self.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	listMenu := self.MenuRepository.FindAll(ctx, tx)
	log.Println(listMenu)
	return ToMenuResponses(listMenu)
}
