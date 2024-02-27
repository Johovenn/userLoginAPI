package main

import (
	"net/http"
	"userloginapi/app"
	"userloginapi/controller"
	"userloginapi/repository"
	"userloginapi/service"
)

func main() {
	db := app.OpenConnection()

	repository := repository.NewUserRepository()
	service := service.NewUserService(repository, db)
	controller := controller.NewUserController(service)

	router := app.NewRouter(controller)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	server.ListenAndServe()
}
