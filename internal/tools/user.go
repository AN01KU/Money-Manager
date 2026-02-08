package tools

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id           uuid.UUID
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	Username     string
	AuthToken    string
}
