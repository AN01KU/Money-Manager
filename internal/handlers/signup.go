package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/AN01KU/money-manager/internal/api"
	"github.com/AN01KU/money-manager/internal/tools"
)

func (h *Handlers) signup(w http.ResponseWriter, r *http.Request) {
	var params api.SignupParams

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		api.RequestErrorHandler(w, errors.New("invalid request body"))
		return
	}

	if params.Username == "" || params.Password == "" || params.Email == "" {
		api.RequestErrorHandler(w, errors.New("username, password, and email are required"))
		return
	}

	if len(params.Password) < 8 {
		api.RequestErrorHandler(w, errors.New("password must be at least 8 characters"))
		return
	}

	database := h.DB
	existingUser := database.GetUserByEmail(params.Email)
	if existingUser != nil {
		http.Error(w, "User with email already exists", http.StatusConflict)
		return
	}

	hashedPassword, err := tools.HashPassword(params.Password)
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	user := database.CreateUser(params.Email, params.Username, hashedPassword)
	if user == nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	jwtToken, err := tools.GenerateJWTToken(user.Id.String())
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	response := api.SignupResponse{
		Token: jwtToken,
		User: api.UserResponse{
			ID:        user.Id.String(),
			Email:     user.Email,
			CreatedAt: user.CreatedAt.Format(time.RFC3339),
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		api.InternalErrorHandler(w)
	}
}
