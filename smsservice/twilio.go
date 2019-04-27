package smsservice

import (
	"net/http"
	"os"
)

var accountSid = os.Getenv("TWILIO_ACCOUNT_SID")
var authToken = os.Getenv("TWILIO_AUTH_TOKEN")

type twilioService struct {
	next   SmsService
	client *http.Client
}

func (twilio *twilioService) Send(sms *SmsMessage) (SmsResult, error) {
	return SmsResult{success: true}, nil
}

func newTwilio(next SmsService) SmsService {
	return &twilioService{
		next:   next,
		client: &http.Client{},
	}
}
