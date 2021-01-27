package services

import (
	"github.com/aipetto/go-aipetto-users-api/src/domain/users"
	"github.com/aipetto/go-aipetto-users-api/src/utils/crypto_utils"
	"github.com/aipetto/go-aipetto-users-api/src/utils/date_utils"
	"github.com/aipetto/go-aipetto-users-api/src/utils/errors"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct {}

type usersServiceInterface interface {
	GetUser(int64) (*users.User, *errors.RestErr)
	CreateUser(users.User) (*users.User, *errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	SearchUserByStatus(string) (users.Users, *errors.RestErr)
	LoginUser(users.LoginRequest) (*users.User, *errors.RestErr)
}

func (s *usersService) GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *usersService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = users.StatusActive
	user.DateCreated = date_utils.GetNowDBFormat()

	hashedPassword, hashingErr := crypto_utils.Hash(user.Password)
	if hashingErr != nil {
		return nil, errors.NewInternalServerError("invalid password")
	}

	user.Password = string(hashedPassword)

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *usersService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr){
	current, err := UsersService.GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if !isPartial{
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.FirstName != "" {
			current.LastName = user.LastName
		}
	}else{
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func (s *usersService) DeleteUser(userId int64) *errors.RestErr{
 	user := &users.User{Id: userId}
 	return user.Delete()
}

func (s *usersService) SearchUserByStatus(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}

func (s *usersService) LoginUser(request users.LoginRequest) (*users.User, *errors.RestErr) {
	dao := &users.User{
		Email: request.Email,
	}
	rawPassword := request.Password

	if err := dao.FindByEmailAndPassword(rawPassword); err != nil {
		return nil, err
	}
	return dao, nil
}