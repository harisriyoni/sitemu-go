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

type beritaControllerImpl struct {
	Service service.BeritaService
}

func NewBeritaController(service service.BeritaService) BeritaController {
	return &beritaControllerImpl{
		Service: service,
	}
}

func (c *beritaControllerImpl) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		helper.WriteError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		helper.WriteError(w, http.StatusBadRequest, "Failed to parse form")
		return
	}

	req := web.BeritaCreateRequest{
		TitleBerita: r.FormValue("title_berita"),
		Tanggal:     r.FormValue("tanggal"),
		Deskripsi:   r.FormValue("deskripsi"),
	}

	imageFile, imageHeader, _ := r.FormFile("image")
	defer func() {
		if imageFile != nil {
			imageFile.Close()
		}
	}()

	response, err := c.Service.Create(r.Context(), userID, req, imageFile, imageHeader)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusCreated, response)
}

func (c *beritaControllerImpl) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		helper.WriteError(w, http.StatusBadRequest, "Failed to parse form")
		return
	}

	req := web.BeritaUpdateRequest{
		TitleBerita: r.FormValue("title_berita"),
		Tanggal:     r.FormValue("tanggal"),
		Deskripsi:   r.FormValue("deskripsi"),
	}

	imageFile, imageHeader, _ := r.FormFile("image")
	defer func() {
		if imageFile != nil {
			imageFile.Close()
		}
	}()

	response, err := c.Service.Update(r.Context(), id, userID, req, imageFile, imageHeader)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, response)
}

func (c *beritaControllerImpl) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

func (c *beritaControllerImpl) GetByUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		helper.WriteError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	results, err := c.Service.GetByUser(r.Context(), userID)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, results)
}

func (c *beritaControllerImpl) GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	results, err := c.Service.GetAll(r.Context())
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, results)
}

func (c *beritaControllerImpl) GetByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	response, err := c.Service.GetByID(r.Context(), id)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, response)
}
