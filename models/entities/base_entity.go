package entities

import (
	"github.com/google/uuid"
	"time"
)

type BaseEntity struct {
	Id        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}
