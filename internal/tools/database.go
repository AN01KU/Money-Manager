package tools

import "github.com/google/uuid"

type DatabaseInterface interface {
	SetupDatabase() error

	CreateUser(email string, username string, passwordHash string) *User
	GetUserByEmail(email string) *User
	GetUserByID(id string) *User

	CreateGroup(name string, createdBy uuid.UUID) *Group
	GetGroupByID(id string) *Group
}

func NewDatabase() (DatabaseInterface, error) {
	var database DatabaseInterface = &mockDB{}

	err := database.SetupDatabase()
	if err != nil {
		return nil, err
	}

	return database, nil
}
