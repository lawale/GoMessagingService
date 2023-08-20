package messageprovider

import (
	"fmt"
	"github.com/lawale/GoMessagingService/models/enums"
)

type PushNotificationMessageProviderService interface {
	SendMessage(model PushNotificationMessageDto) error
}

func GetPushNotificationProviderService(providerType enums.MessageProviderType) (PushNotificationMessageProviderService, error) {
	//TODO: Implement Getting Provider Service
	panic(fmt.Sprintf("Not Implemented For %q", providerType))
}

type PushNotificationMessageDto struct {
	Message      string
	DeviceTokens []string
}

type PushNotificationMessageViewModel struct {
	PushNotificationMessageDto
}
