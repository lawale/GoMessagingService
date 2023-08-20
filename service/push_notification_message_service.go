package service

import (
	"github.com/lawale/GoMessagingService/models/entities"
	"github.com/lawale/GoMessagingService/models/requests"
	"github.com/lawale/GoMessagingService/models/responses"
	"github.com/lawale/GoMessagingService/service/messageprovider"
)

type PushNotificationMessageService struct {
}

func (s *PushNotificationMessageService) SendMessage(request requests.PushNotificationMessageRequest) responses.StatusResponse {
	var provider entities.MessageProvider

	service, err := messageprovider.GetPushNotificationProviderService(provider.ProviderType)

	if err != nil {
		return responses.StatusResponse{
			ResponseCode: responses.ServiceError,
			Message:      "Could Not Send Message",
		}
	}

	err = service.SendMessage(messageprovider.PushNotificationMessageDto{
		Message: request.Content,
	})

	if err != nil {
		return responses.StatusResponse{
			ResponseCode: responses.BadRequest,
			Message:      err.Error(),
		}
	}

	//TODO: Save Message In DB

	return responses.StatusResponse{
		ResponseCode: responses.Success,
		Message:      "Successfully Sent Message",
	}
}
