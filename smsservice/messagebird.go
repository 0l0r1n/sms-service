package smsservice

import (
	"log"

	"github.com/messagebird/go-rest-api"
	"github.com/messagebird/go-rest-api/sms"
)

type messageBirdService struct {
	next   *SmsService
	client *messagebird.Client
}

// NewMessageBird initiates the messagebird service is a gateway for sending messages using messagebird
func NewMessageBird(next *SmsService) SmsService {
	client := messagebird.New("") // TODO: insert key here
	mb := &messageBirdService{
		next:   next,
		client: client,
	}
	return mb
}

func (mb *messageBirdService) send(message SmsMessage) (SmsResult, error) {
	msg, err := sms.Create(mb.client, "+31XXXXXXXXX",
		[]string{"+31XXXXXXXXX"},
		"Hi! This is your first message",
		nil)
	if err != nil {
		log.Println(err)
	}
	log.Println(msg)
	return SmsResult{success: true}, nil
}
