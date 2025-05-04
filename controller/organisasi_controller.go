package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type OrganisasiController interface {
	Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	GetByUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	GetAll(w http.ResponseWriter, r *http.Request, ps httprouter.Params)

}
