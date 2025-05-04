package controller

import (
	"encoding/json"
	"net/http"

	"github.com/harisriyoni/sitemu-go/middleware"
	"github.com/harisriyoni/sitemu-go/model/web"
	"github.com/harisriyoni/sitemu-go/service"
	"github.com/julienschmidt/httprouter"
)

type userControllerImpl struct {
	UserService service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &userControllerImpl{UserService: service}
}

func (c *userControllerImpl) Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var req web.UserRegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	res, err := c.UserService.Register(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(res)
}

func (c *userControllerImpl) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var req web.UserLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	token, expiresIn, err := c.UserService.Login(r.Context(), req)
	if err != nil {
		http.Error(w, "Login failed: Your Username or Password is Incorrect!", http.StatusUnauthorized)
		return
	}

	response := map[string]interface{}{
		"message":      "Berhasil login",
		"token":        token,
		"expires_in":   int(expiresIn),
		"expires_unit": "hours",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (c *userControllerImpl) Profile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	res, err := c.UserService.GetProfile(r.Context(), userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(res)
}

func (c *userControllerImpl) Update(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Ambil user ID dari JWT (context)
	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Decode body request
	var req web.UserUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON input", http.StatusBadRequest)
		return
	}

	// Lakukan update via service
	res, err := c.UserService.UpdateProfile(r.Context(), userID, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Berhasil
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (c *userControllerImpl) Delete(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if err := c.UserService.DeleteAccount(r.Context(), userID); err != nil {
		http.Error(w, "Delete failed", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"message": "Account successfully deleted",
	}
	json.NewEncoder(w).Encode(response)
}
