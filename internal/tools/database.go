package tools

import (
	"time"

	"github.com/google/uuid"
)

type UserDetails struct {
	Id           uuid.UUID
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	Username     string
	AuthToken    string
}

type DatabaseInterface interface {
	CreateUser(email string, username string, passwordHash string) (*UserDetails, error)
	GetUserByEmail(email string) *UserDetails
	GetUserByID(id string) *UserDetails
	SetupDatabase() error
}

func NewDatabase() (*DatabaseInterface, error) {
	var database DatabaseInterface = &mockDB{}

	var err error = database.SetupDatabase()
	if err != nil {
		return nil, err
	}

	return &database, nil
}
