package entities

import (
	"github.com/google/uuid"
)

type MessageDeliveryLog struct {
	BaseEntity
	MessagingId uuid.UUID
}
