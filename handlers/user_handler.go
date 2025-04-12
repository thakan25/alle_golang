package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"task-manager/common"
	"task-manager/models/dtos"
	"task-manager/service"
)

// UserHandler handles HTTP requests for users
type UserHandler struct {
	service *service.UserService
}

// NewUserHandler creates a new user handler
func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

// respondJSON writes a JSON response with the given status code
func (h *UserHandler) respondJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

// HandleUsers handles requests to /users
func (h *UserHandler) HandleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.CreateUser(w, r)
	case http.MethodGet:
		h.GetUsers(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// HandleUser handles requests to /users/{id}
func (h *UserHandler) HandleUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetUser(w, r)
	case http.MethodDelete:
		h.DeleteUser(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// CreateUser handles POST /users
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var dto dtos.CreateUserDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		panic(common.ErrInvalidRequest)
	}

	user, err := h.service.CreateUser(dto)
	if err != nil {
		panic(err)
	}

	h.respondJSON(w, http.StatusCreated, user)
}

// GetUsers handles GET /users
func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetUsers()
	if err != nil {
		panic(err)
	}

	h.respondJSON(w, http.StatusOK, users)
}

// GetUser handles GET /users/{id}
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from URL
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	id := parts[2]

	user, err := h.service.GetUser(id)
	if err != nil {
		panic(err)
	}

	h.respondJSON(w, http.StatusOK, user)
}

// DeleteUser handles DELETE /users/{id}
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from URL
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	id := parts[2]

	if err := h.service.DeleteUser(id); err != nil {
		panic(err)
	}

	h.respondJSON(w, http.StatusNoContent, nil)
} 