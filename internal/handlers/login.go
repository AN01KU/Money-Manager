package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/AN01KU/money-manager/internal/api"
	"github.com/AN01KU/money-manager/internal/tools"
)

func (h *Handlers) login(w http.ResponseWriter, r *http.Request) {
	var params api.LoginParams

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		api.RequestErrorHandler(w, err)
		return
	}

	if params.Email == "" || params.Password == "" {
		http.Error(w, "Empty email or password", http.StatusBadRequest)
		return
	}

	database := h.DB
	user := database.GetUserByEmail(params.Email)
	if user == nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	if !tools.CheckPasswordHash(params.Password, user.PasswordHash) {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
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
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		api.InternalErrorHandler(w)
		return
	}
}
