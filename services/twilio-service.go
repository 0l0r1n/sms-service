package main

// TwilioService is a gateway for sending messages using twilio
type TwilioService struct {
	next SmsService
}
