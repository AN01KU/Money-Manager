package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/AN01KU/money-manager/internal/api"
	"github.com/AN01KU/money-manager/internal/tools"
	"github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"
)

func (h *Handlers) signup(w http.ResponseWriter, r *http.Request) {
	var params = api.SignupParams{}

	var err error
	err = json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		api.RequestErrorHandler(w, errors.New("Invalid request body"))
		return
	}

	// Validate the params
	username := params.Username
	password := params.Password
	email := params.Email
	if len(username) == 0 || len(password) == 0 || len(email) == 0 {
		api.RequestErrorHandler(w, errors.New("Empty username or password or email"))
		return
	}

	// check if user with email already exists or not
	database := h.DB
	existingUser := (*database).GetUserByEmail(email)
	if existingUser != nil {
		//TODO: not sure which error to throwdon
		api.RequestErrorHandler(w, errors.New("User already exists"))
		return
	}

	// create user
	hashedPassword, err := tools.HashPassword(password)
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}
	userdetails, err := (*database).CreateUser(email, username, hashedPassword)
	if err != nil {
		//TODO: not sure which error to throw
		api.InternalErrorHandler(w)
		return
	}

	// generate jwt token
	claims := jwt.MapClaims{
		"user_id": userdetails.Id.String(),
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	response := api.SignupResponse{
		Token: tokenString,
		User: api.UserResponse{
			ID:        userdetails.Id.String(),
			Email:     userdetails.Email,
			CreatedAt: userdetails.CreatedAt.Format(time.RFC3339),
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

	log.Info("User registered successfully")

}
