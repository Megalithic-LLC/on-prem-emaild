package cloudservice

import (
	"github.com/on-prem-net/emaild/cloudservice/emailproto"
)

func (self *CloudService) SendErrorResponse(requestId uint64, err error) error {
	errorRes := emailproto.ClientMessage{
		Id: requestId,
		MessageType: &emailproto.ClientMessage_ErrorResponse{
			ErrorResponse: &emailproto.ErrorResponse{
				Error: err.Error(),
			},
		},
	}
	return self.SendResponse(errorRes)
}
