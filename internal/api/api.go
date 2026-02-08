package api

import (
	"encoding/json"
	"net/http"
)

type SignupParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type SignupResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}
type LoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

type CreateGroupParams struct {
	Name string `json:"name"`
}

type GroupResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedBy string `json:"created_by"`
	CreatedAt string `json:"created_at"`
}

// type GroupResponse struct{ ID, Name, CreatedBy, CreatedAt string }

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}
