package repository

import (
	"go-react-todo/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}

type userRepository struct {
	dbConn *gorm.DB
}

func NewUserRepository(dbConn *gorm.DB) IUserRepository {
	return &userRepository{dbConn}
}

func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
	if err := ur.dbConn.Where("email = ?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) CreateUser(user *model.User) error {
	if err := ur.dbConn.Create(user).Error; err != nil {
		return err
	}
	return nil
}
