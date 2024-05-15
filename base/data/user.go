package data

import (
	"go-video/base/model"

	"gorm.io/gorm"
)

type UserData interface {
	GetUser(id int) (*model.User, error)
	CreateUser(user *model.User) error
	GetUserByPhone(phone string) (*model.User, error)
}
type UserDataImpl struct {
	DB *gorm.DB
}

func NewUserData(db *gorm.DB) UserData {
	return &UserDataImpl{DB: db}
}

func (d *UserDataImpl) GetUser(id int) (*model.User, error) {
	var user model.User
	result := d.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (d *UserDataImpl) CreateUser(user *model.User) error {
	result := d.DB.Create(user)
	return result.Error
}

func (d *UserDataImpl) GetUserByPhone(phone string) (*model.User, error) {
	user := &model.User{}
	result := d.DB.Where("phone = ?", phone).First(user)
	return user, result.Error
}
