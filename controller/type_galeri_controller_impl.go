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

type typeGaleriControllerImpl struct {
	Service service.TypeGaleriService
}

func NewTypeGaleriController(service service.TypeGaleriService) TypeGaleriController {
	return &typeGaleriControllerImpl{
		Service: service,
	}
}

// ✅ Create (Protected)
func (c *typeGaleriControllerImpl) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		helper.WriteError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	request := web.TypeGaleriCreateRequest{
		Type: r.FormValue("type"),
	}

	result, err := c.Service.Create(r.Context(), request, userID)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusCreated, result)
}

// ✅ Update (Protected)
func (c *typeGaleriControllerImpl) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	request := web.TypeGaleriUpdateRequest{
		Type: r.FormValue("type"),
	}

	result, err := c.Service.Update(r.Context(), id, userID, request)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, result)
}

// ✅ Delete (Protected)
func (c *typeGaleriControllerImpl) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

// ✅ Get All (Public)
func (c *typeGaleriControllerImpl) GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	results, err := c.Service.GetAll(r.Context())
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, results)
}

// ✅ Get By ID (Public)
func (c *typeGaleriControllerImpl) GetByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

// ✅ Get By User (Protected)
func (c *typeGaleriControllerImpl) GetByUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		helper.WriteError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	results, err := c.Service.GetByUser(r.Context(), userID)
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, results)
}
