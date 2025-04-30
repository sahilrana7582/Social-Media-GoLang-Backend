package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sahilrana7582/Social-Media-GoLang-Backend/internal/model"
	"github.com/sahilrana7582/Social-Media-GoLang-Backend/internal/service"
)

type UserHandler struct {
	UserService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	var user model.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	createdUser, err := h.UserService.CreateUser(user.Name, user.Email)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}
