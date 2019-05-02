package smsservice

import (
	"log"
	"os"

	"github.com/messagebird/go-rest-api"
	"github.com/messagebird/go-rest-api/sms"
)

type messageBirdService struct {
	next   SmsService
	client *messagebird.Client
}

func newMessageBird(next SmsService) SmsService {
	client := messagebird.New(os.Getenv("MESSAGE_BIRD_KEY"))
	mb := &messageBirdService{
		next:   next,
		client: client,
	}
	return mb
}

func (mb *messageBirdService) Send(message *SmsMessage) SmsResult {
	_, err := sms.Create(mb.client, message.Originator, []string{message.Recipient}, message.Body, nil)
	if err != nil {
		switch errResp := err.(type) {
		case messagebird.ErrorResponse:
			for _, mbError := range errResp.Errors {
				log.Printf("Error: %#v\n", mbError)
			}
		}
		if mb.next == nil {
			return errorResult()
		}
		return mb.next.Send(message)
	}

	return SmsResult{Success: true}
}
