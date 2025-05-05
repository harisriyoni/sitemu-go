package controller

import (
	"net/http"
	"strconv"

	"github.com/harisriyoni/sitemu-go/helper"
	"github.com/harisriyoni/sitemu-go/middleware"
	"github.com/harisriyoni/sitemu-go/model/web"
	"github.com/harisriyoni/sitemu-go/service"
	"github.com/julienschmidt/httprouter"
)

type prestasiControllerImpl struct {
	Service service.PrestasiService
}

func NewPrestasiController(service service.PrestasiService) PrestasiController {
	return &prestasiControllerImpl{Service: service}
}

func (c *prestasiControllerImpl) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		helper.WriteError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	req := web.PrestasiCreateRequest{
		Title:     r.FormValue("title"),
		Tahun:     r.FormValue("tahun"),
		Prestasi:  r.FormValue("prestasi"),
		Deskripsi: r.FormValue("deskripsi"),
	}

	result, err := c.Service.Create(r.Context(), userID, req)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusCreated, result)
}

func (c *prestasiControllerImpl) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		helper.WriteError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	req := web.PrestasiUpdateRequest{
		Title:     r.FormValue("title"),
		Tahun:     r.FormValue("tahun"),
		Prestasi:  r.FormValue("prestasi"),
		Deskripsi: r.FormValue("deskripsi"),
	}

	result, err := c.Service.Update(r.Context(), id, userID, req)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, result)
}

func (c *prestasiControllerImpl) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		helper.WriteError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	err = c.Service.Delete(r.Context(), id, userID)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, map[string]string{"message": "Deleted successfully"})
}

func (c *prestasiControllerImpl) GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	list, err := c.Service.GetAll(r.Context())
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	helper.WriteJSON(w, http.StatusOK, list)
}

func (c *prestasiControllerImpl) GetByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	result, err := c.Service.GetByID(r.Context(), id)
	if err != nil {
		helper.WriteError(w, http.StatusNotFound, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, result)
}

func (c *prestasiControllerImpl) GetByUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		helper.WriteError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	list, err := c.Service.GetByUser(r.Context(), userID)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, list)
}
