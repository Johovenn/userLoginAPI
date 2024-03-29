package service

import "userloginapi/models/domain"

type UserService interface {
	Login(user domain.User) (domain.User, error, bool)
	Register(user domain.User) (domain.User, error)
}
