package cafe

import(
	"github.com/go-playground/validator/v10"
	"database/sql"
	"context"
	"log"

	"github.com/michaelact/cafe-everywhere/exception"
	"github.com/michaelact/cafe-everywhere/helper"
)

type CafeService interface {
	Login(ctx context.Context, request HTTPCafeRequest) (HTTPCafeResponse, bool)
	Create(ctx context.Context, request HTTPCafeRequest) HTTPCafeResponse
	Update(ctx context.Context, request HTTPCafeUpdateRequest) HTTPCafeResponse
	Delete(ctx context.Context, id int)
	FindById(ctx context.Context, id int) HTTPCafeResponse
	FindByEmail(ctx context.Context, email string) HTTPCafeResponse
	FindAll(ctx context.Context) []HTTPCafeResponse
}

type CafeServiceImpl struct {
	CafeRepository CafeRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewCafeService(CafeRepository CafeRepository, db *sql.DB, validate *validator.Validate) CafeService {
	return &CafeServiceImpl{
		CafeRepository: CafeRepository, 
		DB:             db, 
		Validate:       validate, 
	}
}

func (self *CafeServiceImpl) Login(ctx context.Context, request HTTPCafeRequest) (HTTPCafeResponse, bool) {
	tx, err := self.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Fail if existing cafe not found
	cafe, err := self.CafeRepository.FindByEmail(ctx, tx, request.Email)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	// Verify cafe authentication
	ok := helper.VerifyPassword(cafe.Password, request.Password)
	if !ok {
		return HTTPCafeResponse{}, ok
	} 

	return ToCafeResponse(cafe), ok
}

func (self *CafeServiceImpl) Create(ctx context.Context, request HTTPCafeRequest) HTTPCafeResponse {
	// Fail if request is not valid
	err := self.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := self.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	passwordHash, err := helper.GetPasswordHash(request.Password)
	helper.PanicIfError(err)

	cafe := CafeDatabaseIO{
		Email:    request.Email,
		Title:    request.Title,
		Password: passwordHash,
	}

	cafe = self.CafeRepository.Insert(ctx, tx, cafe)
	return ToCafeResponse(cafe)
}

func (self *CafeServiceImpl) Update(ctx context.Context, request HTTPCafeUpdateRequest) HTTPCafeResponse {
	// Fail if request is not valid
	err := self.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := self.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Fail if existing cafe not found
	cafe, err := self.CafeRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	// Update existing cafe
	cafe.Email = request.Email
	cafe.Title = request.Title
	cafe.Password = request.Password
	cafe = self.CafeRepository.Update(ctx, tx, cafe)
	return ToCafeResponse(cafe)
}

func (self *CafeServiceImpl) Delete(ctx context.Context, id int) {
	tx, err := self.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Fail if existing cafe not found
	_, err = self.CafeRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	// Delete existing cafe
	self.CafeRepository.Delete(ctx, tx, id)
}

func (self *CafeServiceImpl) FindByEmail(ctx context.Context, email string) HTTPCafeResponse {
	tx, err := self.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Fail if existing cafe not found
	cafe, err := self.CafeRepository.FindByEmail(ctx, tx, email)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return ToCafeResponse(cafe)
}

func (self *CafeServiceImpl) FindById(ctx context.Context, id int) HTTPCafeResponse {
	tx, err := self.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Fail if existing cafe not found
	cafe, err := self.CafeRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return ToCafeResponse(cafe)
}

func (self *CafeServiceImpl) FindAll(ctx context.Context) []HTTPCafeResponse {
	tx, err := self.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	listCafe := self.CafeRepository.FindAll(ctx, tx)
	log.Println(listCafe)
	return ToCafeResponses(listCafe)
}