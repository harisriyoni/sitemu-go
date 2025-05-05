package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type BeritaController interface {
	Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	GetAll(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	GetByUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	GetByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}
