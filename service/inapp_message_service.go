package service

import (
	"github.com/lawale/GoMessagingService/models/requests"
	"github.com/lawale/GoMessagingService/models/responses"
)

type InAppMessageService struct {
}

func (s *InAppMessageService) SendMessage(request requests.InAppMessageRequest) responses.StatusResponse {

	//TODO: Save Message In DB

	return responses.StatusResponse{
		ResponseCode: responses.Success,
		Message:      "Successfully Sent Message",
	}
}
