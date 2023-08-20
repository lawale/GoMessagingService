package entities

import "github.com/lawale/GoMessagingService/models/enums"

type Message struct {
	BaseEntity
	UserId  string
	Content string
	Status  enums.MessageDeliveryStatus
}
