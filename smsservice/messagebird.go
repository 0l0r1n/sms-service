package smsservice

import (
	"fmt"
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

func (mb *messageBirdService) Send(message *SmsMessage) (SmsResult, error) {
	msg, err := sms.Create(mb.client, message.Originator, []string{message.Recipient}, message.Body, nil)
	if err != nil {
		switch errResp := err.(type) {
		case messagebird.ErrorResponse:
			for _, mbError := range errResp.Errors {
				fmt.Printf("Error: %#v\n", mbError)
			}
		}
		if mb.next == nil {
			return errorResult(err)
		}
		return mb.next.Send(message)
	}

	log.Println(msg)
	return SmsResult{success: true}, nil
}
