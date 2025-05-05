package controller

import (
	"net/http"
	"strconv"

	"github.com/harisriyoni/sitemu-go/helper"
	"github.com/harisriyoni/sitemu-go/model/web"
	"github.com/harisriyoni/sitemu-go/service"
	"github.com/julienschmidt/httprouter"
)

type galeriControllerImpl struct {
	Service service.GaleriService
}

func NewGaleriController(service service.GaleriService) GaleriController {
	return &galeriControllerImpl{Service: service}
}

func (c *galeriControllerImpl) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		helper.WriteError(w, http.StatusBadRequest, "Failed to parse form")
		return
	}

	req := web.GaleriCreateRequest{
		TitleImage:   r.FormValue("title_image"),
		TypeGaleriID: helper.Atoi(r.FormValue("type_galeri_id")),
	}

	imageFile, imageHeader, _ := r.FormFile("image")
	defer func() {
		if imageFile != nil {
			imageFile.Close()
		}
	}()

	result, err := c.Service.Create(r.Context(), req, imageFile, imageHeader)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusCreated, result)
}

func (c *galeriControllerImpl) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		helper.WriteError(w, http.StatusBadRequest, "Failed to parse form")
		return
	}

	req := web.GaleriUpdateRequest{
		TitleImage:   r.FormValue("title_image"),
		TypeGaleriID: helper.Atoi(r.FormValue("type_galeri_id")),
	}

	imageFile, imageHeader, _ := r.FormFile("image")
	defer func() {
		if imageFile != nil {
			imageFile.Close()
		}
	}()

	result, err := c.Service.Update(r.Context(), id, req, imageFile, imageHeader)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, result)
}

func (c *galeriControllerImpl) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	err = c.Service.Delete(r.Context(), id)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, map[string]string{"message": "Deleted successfully"})
}

func (c *galeriControllerImpl) GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	list, err := c.Service.GetAll(r.Context())
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, list)
}

func (c *galeriControllerImpl) GetByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
