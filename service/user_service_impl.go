package service

import (
	"errors"
	"userloginapi/helper"
	"userloginapi/models/domain"
	"userloginapi/repository"

	"gorm.io/gorm"
)

type UserServiceImpl struct {
	repository repository.UserRepository
	DB         *gorm.DB
}

func NewUserService(repository repository.UserRepository, db *gorm.DB) UserService{
	return &UserServiceImpl{
		repository:  repository,
		DB: db,
	}
}

func (u *UserServiceImpl) Login(user domain.User) (domain.User , error){
	tx := u.DB.Begin()

	foundUser, err := u.repository.Login(tx, user)
	helper.PanicIfError(err)

	if foundUser.Password == user.Password{
		tx.Commit()
		return foundUser, nil
	} else{
		tx.Rollback()
		return nil, errors.New("Invalid Login Credentials")
	}
}