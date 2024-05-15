package service

import (
	"fmt"
	"go-video/base/data"
	"go-video/base/dto"
	"go-video/base/model"
)

type UserService interface {
	GetUser(id int) (*model.User, error)
	CreateUser(user *dto.CreateUserRequest) error
	LoginUser(user *dto.LoginUserRequest, tokenString string) error
}

type UserServiceImpl struct {
	UserData  data.UserData
	JWTSecret string
}

func NewUserService(userData data.UserData, jwtSecret string) UserService {
	return &UserServiceImpl{UserData: userData, JWTSecret: jwtSecret}
}

func (u UserServiceImpl) GetUser(id int) (*model.User, error) {
	user, err := u.UserData.GetUser(id)
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:    user.ID,
		Phone: user.Phone,
		Email: user.Email,
	}, nil
}

func (u UserServiceImpl) CreateUser(createUserRequest *dto.CreateUserRequest) error {
	if createUserRequest.Phone == "" || createUserRequest.Password == "" {
		return fmt.Errorf("the phone or password is required")
	}

	_, err := u.UserData.GetUserByPhone(createUserRequest.Phone)
	if err == nil {
		return fmt.Errorf("手机号 %s 已被注册", createUserRequest.Phone)
	}

	return u.UserData.CreateUser(&model.User{
		Phone:    createUserRequest.Phone,
		Password: createUserRequest.Password,
	})
}

func (u UserServiceImpl) LoginUser(loginUserRequest *dto.LoginUserRequest, tokenString string) error {

	return nil
}
