package user

import(
	"github.com/go-playground/validator/v10"
	"database/sql"
	"context"
	"log"

	"github.com/michaelact/cafe-everywhere/exception"
	"github.com/michaelact/cafe-everywhere/helper"
)

type UserService interface {
	Login(ctx context.Context, request HTTPUserRequest) (HTTPUserResponse, bool)
	Create(ctx context.Context, request HTTPUserRequest) HTTPUserResponse
	Update(ctx context.Context, request HTTPUserUpdateRequest) HTTPUserResponse
	Delete(ctx context.Context, id int)
	FindById(ctx context.Context, id int) HTTPUserResponse
	FindByEmail(ctx context.Context, email string) HTTPUserResponse
	FindAll(ctx context.Context) []HTTPUserResponse
}

type UserServiceImpl struct {
	UserRepository UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(UserRepository UserRepository, db *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: UserRepository, 
		DB:                 db, 
		Validate:           validate, 
	}
}

func (self *UserServiceImpl) Login(ctx context.Context, request HTTPUserRequest) (HTTPUserResponse, bool) {
	tx, err := self.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Fail if existing user not found
	user, err := self.UserRepository.FindByEmail(ctx, tx, request.Email)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	// Verify user authentication
	ok := helper.VerifyPassword(user.Password, request.Password)
	if !ok {
		return HTTPUserResponse{}, ok
	} 

	return ToUserResponse(user), ok
}

func (self *UserServiceImpl) Create(ctx context.Context, request HTTPUserRequest) HTTPUserResponse {
	// Fail if request is not valid
	err := self.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := self.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	passwordHash, err := helper.GetPasswordHash(request.Password)
	helper.PanicIfError(err)

	user := UserDatabaseIO{
		Email:    request.Email,
		Password: passwordHash,
	}

	user = self.UserRepository.Insert(ctx, tx, user)
	return ToUserResponse(user)
}

func (self *UserServiceImpl) Update(ctx context.Context, request HTTPUserUpdateRequest) HTTPUserResponse {
	// Fail if request is not valid
	err := self.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := self.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Fail if existing user not found
	user, err := self.UserRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	// Update existing user
	user.Email = request.Email
	user.Password = request.Password
	user = self.UserRepository.Update(ctx, tx, user)
	return ToUserResponse(user)
}

func (self *UserServiceImpl) Delete(ctx context.Context, id int) {
	tx, err := self.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Fail if existing user not found
	_, err = self.UserRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	// Delete existing user
	self.UserRepository.Delete(ctx, tx, id)
}

func (self *UserServiceImpl) FindByEmail(ctx context.Context, email string) HTTPUserResponse {
	tx, err := self.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Fail if existing user not found
	user, err := self.UserRepository.FindByEmail(ctx, tx, email)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return ToUserResponse(user)
}

func (self *UserServiceImpl) FindById(ctx context.Context, id int) HTTPUserResponse {
	tx, err := self.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Fail if existing user not found
	user, err := self.UserRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return ToUserResponse(user)
}

func (self *UserServiceImpl) FindAll(ctx context.Context) []HTTPUserResponse {
	tx, err := self.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	listUser := self.UserRepository.FindAll(ctx, tx)
	log.Println(listUser)
	return ToUserResponses(listUser)
}