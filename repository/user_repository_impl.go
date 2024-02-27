package repository

import (
	"userloginapi/models/domain"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct{}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (u *UserRepositoryImpl) Login(db *gorm.DB, user domain.User) (domain.User, error) {
	username := user.Username
	var foundUser domain.User

	err := db.Where("username = ?", username).First(&foundUser).Error

	return foundUser, err
}
