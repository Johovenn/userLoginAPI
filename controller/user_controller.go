package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserController interface {
	Login(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}
