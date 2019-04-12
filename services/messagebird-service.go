package main

// MessageBirdService is a gateway for sending messages using twilio
type MessageBirdService struct {
	next SmsService
}
