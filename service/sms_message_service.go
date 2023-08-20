package service

import (
	"github.com/lawale/GoMessagingService/models/entities"
	"github.com/lawale/GoMessagingService/models/requests"
	"github.com/lawale/GoMessagingService/models/responses"
	"github.com/lawale/GoMessagingService/service/messageprovider"
)

type SmsMessageService struct {
}

func (s *SmsMessageService) SendMessage(request requests.SmsMessageRequest) responses.StatusResponse {
	var provider entities.MessageProvider

	service, err := messageprovider.GetSmsProviderService(provider.ProviderType)

	if err != nil {
		return responses.StatusResponse{
			ResponseCode: responses.ServiceError,
			Message:      "Could Not Send Message",
		}
	}

	err = service.SendMessage(messageprovider.SmsMessageDto{
		Message:     request.Content,
		PhoneNumber: request.PhoneNumber,
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
