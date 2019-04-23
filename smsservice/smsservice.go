package smsservice

// SmsMessage is a sms message that will be handled by a SmsService
type SmsMessage struct {
	Originator string
	Recipients []string
	Body       string
}

// SmsService represents a Sms notification gateway
type SmsService interface {
	Send(sms SmsMessage) (SmsResult, error)
}

// SmsResult is returned when a Sms request is sent
type SmsResult struct {
	success bool
}

// GetChainOfServices returns the chain of responsibility, starting by Twilio, then message bird and lastly simple notification service
func GetChainOfServices() SmsService {
	return &twilioService{next: newMessageBird(nil)}
}
