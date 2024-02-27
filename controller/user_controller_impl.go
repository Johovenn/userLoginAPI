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

	user, err := u.service.Login(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Internal Server Error")
	}

	response := response.Response{
		Code:   200,
		Status: "Success Login",
	}

	helper.WriteResponse(w, response)
}
