package repository

import (
	"userloginapi/models/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	Login(db *gorm.DB, user domain.User) (domain.User, error)
	Register(db *gorm.DB, user domain.User) (domain.User, error)
}