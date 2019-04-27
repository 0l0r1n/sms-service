package smsservice

// SmsMessage is a sms message that will be handled by a SmsService
type SmsMessage struct {
	Originator string
	Recipients []string
	Body       string
}

// SmsService represents a Sms notification gateway
type SmsService interface {
	Send(sms *SmsMessage) (SmsResult, error)
}

// SmsResult is returned when a Sms request is sent
type SmsResult struct {
	success bool
}

// GetChainOfServices returns the chain of responsibility with fallback implementation for sending SMSs
func GetChainOfServices() SmsService {
	return newMessageBird(nil)
}

func returnErrorResult(err error) (SmsResult, error) {
	return SmsResult{success: false}, err
}
