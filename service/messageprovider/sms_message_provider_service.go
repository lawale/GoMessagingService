package messageprovider

import (
	"fmt"
	"github.com/lawale/GoMessagingService/models/enums"
)

type SmsMessageProviderService interface {
	SendMessage(model SmsMessageDto) error
}

func GetSmsProviderService(providerType enums.MessageProviderType) (SmsMessageProviderService, error) {
	//TODO: Implement Getting Provider Service
	panic(fmt.Sprintf("Not Implemented For %q", providerType))
}

type SmsMessageDto struct {
	Message     string
	PhoneNumber string
}

type SmsMessageViewModel struct {
	SmsMessageDto
}
