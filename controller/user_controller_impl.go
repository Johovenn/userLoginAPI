package controller

import (
	"fmt"
	"net/http"
	"userloginapi/helper"
	"userloginapi/models/domain"
	"userloginapi/models/response"
	"userloginapi/service"

	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &UserControllerImpl{
		service: service,
	}
}

func (u *UserControllerImpl) Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var user domain.User
	helper.ReadFromRequestBody(r, &user)

	user, err, success := u.service.Login(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Internal Server Error")
	}

	var res response.LoginResponse

	if success {
		res.Code = 200
		res.Status = "Success Login"
		res.Data = user
	} else {
		res.Code = 401
		res.Status = "Invalid username or password"
		res.Data = nil
	}

	helper.WriteResponse(w, res)
}
