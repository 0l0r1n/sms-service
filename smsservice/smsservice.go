package smsservice

// SmsMessage is a sms message that will be handled by a SmsService
type SmsMessage struct {
	Originator string `json:"originator"`
	Recipient  string `json:"recipient"`
	Body       string `json:"body"`
}

// SmsService represents a Sms notification gateway
type SmsService interface {
	Send(message *SmsMessage) SmsResult
}

// SmsResult is returned when a Sms request is sent
type SmsResult struct {
	Success bool `json:"success"`
}

// GetChainOfServices returns the chain of responsibility with fallback implementation for sending SMSs
func GetChainOfServices() SmsService {
	return newMessageBird(newTwilio(nil))
}

func errorResult() SmsResult {
	return SmsResult{Success: false}
}
