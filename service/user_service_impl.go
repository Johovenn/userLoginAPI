package service

import (
	"userloginapi/helper"
	"userloginapi/models/domain"
	"userloginapi/repository"

	"gorm.io/gorm"
)

type UserServiceImpl struct {
	repository repository.UserRepository
	DB         *gorm.DB
}

func NewUserService(repository repository.UserRepository, db *gorm.DB) UserService {
	return &UserServiceImpl{
		repository: repository,
		DB:         db,
	}
}

func (u *UserServiceImpl) Login(user domain.User) (domain.User, error, bool) {
	tx := u.DB.Begin()

	foundUser, err := u.repository.Login(tx, user)
	helper.PanicIfError(err)

	if foundUser.Password == user.Password {
		tx.Commit()
		return foundUser, nil, true
	} else {
		tx.Rollback()
		return domain.User{}, err, false
	}
}

func (u *UserServiceImpl) Register(user domain.User) (domain.User, error) {
	tx := u.DB.Begin()

	result, err := u.repository.Login(tx, user)
	helper.RollbackOrCommit(tx, err)

	if err != nil {
		tx.Rollback()
		return result, err
	} else {
		tx.Commit()
		return result, err
	}
}
