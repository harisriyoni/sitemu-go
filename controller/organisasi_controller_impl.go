package controller

import (
	"encoding/json"
	"net/http"

	"github.com/harisriyoni/sitemu-go/helper"
	"github.com/harisriyoni/sitemu-go/middleware"
	"github.com/harisriyoni/sitemu-go/model/web"
	"github.com/harisriyoni/sitemu-go/service"
	"github.com/julienschmidt/httprouter"
)

type organisasiControllerImpl struct {
	Service service.OrganisasiService
}

func NewOrganisasiController(service service.OrganisasiService) OrganisasiController {
	return &organisasiControllerImpl{Service: service}
}

func (c *organisasiControllerImpl) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	r.ParseMultipartForm(10 << 20)
	image, imageHeader, _ := r.FormFile("image")

	req := web.OrganisasiCreateRequest{
		Jabatan: r.FormValue("jabatan"),
		Nama:    r.FormValue("nama"),
	}

	res, err := c.Service.CreateOrganisasi(r.Context(), userID, req, image, imageHeader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(res)
}

func (c *organisasiControllerImpl) GetByUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	res, err := c.Service.GetOrganisasiByUserID(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(res)
}

func (c *organisasiControllerImpl) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	r.ParseMultipartForm(10 << 20)
	image, imageHeader, _ := r.FormFile("image")

	req := web.OrganisasiUpdateRequest{
		Jabatan: r.FormValue("jabatan"),
		Nama:    r.FormValue("nama"),
	}

	id := ps.ByName("id")
	intID, err := helper.StringToInt(id)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}
	res, err := c.Service.UpdateOrganisasi(r.Context(), intID, userID, req, image, imageHeader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(res)
}

func (c *organisasiControllerImpl) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	id := ps.ByName("id")
	intID, err := helper.StringToInt(id)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	err = c.Service.DeleteOrganisasi(r.Context(), intID, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Deleted successfully"})
}

func (c *organisasiControllerImpl) GetAll(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()

	result, err := c.Service.GetAllOrganisasi(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}
