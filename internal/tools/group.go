package tools

import (
	"time"

	"github.com/google/uuid"
)

type Group struct {
	Id        uuid.UUID
	Name      string
	CreatedBy uuid.UUID
	CreatedAt time.Time
}
