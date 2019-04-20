package smsservice

// Sms is a sms message that will be handled by the SmsService
type SmsMessage struct {
	originator string
}

// SmsService represents a Sms notification gateway
type SmsService interface {
	send(sms SmsMessage) (SmsResult, error)
}

// SmsResult is returned when a Sms request is sent
type SmsResult struct {
	success bool
}

// GetChainOfServices returns the chain of responsibility, starting by Twilio, then message bird and lastly simple notification service
func GetChainOfServices() SmsService {
	return &TwilioService{next: NewMessageBird(nil)}
}
