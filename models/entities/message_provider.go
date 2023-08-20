package entities

import "github.com/lawale/GoMessagingService/models/enums"

type MessageProvider struct {
	BaseEntity
	Name         string
	MessageTypes enums.MessageType
	ProviderType enums.MessageProviderType
	IsActive     bool
}
