package smsservice

// Sms is a sms message that will be handled by the SmsService
type Sms struct {
	originator string
}

// SmsService represents a Sms notification gateway
type SmsService interface {
	send(sms Sms)
}

// GetChainOfServices returns the chain of responsibility, starting by Twilio, then message bird and lastly simple notification service
func GetChainOfServices() SmsService {
	return &TwilioService{&MessageBirdService{&SimpleNotificationService{}}}
}
