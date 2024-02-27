package app

import (
	"userloginapi/controller"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(controller controller.UserController) *httprouter.Router{
	router := httprouter.New()

	router.POST("/login", controller.Login)

	return router
}
