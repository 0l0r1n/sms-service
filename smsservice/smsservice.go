package smsservice

// SmsMessage is a sms message that will be handled by a SmsService
type SmsMessage struct {
	Originator string
	Recipient  string
	Body       string
}

// SmsService represents a Sms notification gateway
type SmsService interface {
	Send(message *SmsMessage) (SmsResult, error)
}

// SmsResult is returned when a Sms request is sent
type SmsResult struct {
	success bool
}

// GetChainOfServices returns the chain of responsibility with fallback implementation for sending SMSs
func GetChainOfServices() SmsService {
	return newMessageBird(newTwilio(nil))
}

func errorResult(err error) (SmsResult, error) {
	return SmsResult{success: false}, err
}
