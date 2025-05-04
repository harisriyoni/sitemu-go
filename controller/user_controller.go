package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserController interface {
	Register(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	Profile(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}
