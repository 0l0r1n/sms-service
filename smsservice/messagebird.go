package smsservice

// MessageBirdService is a gateway for sending messages using twilio
type MessageBirdService struct {
	next SmsService
}

func (mb MessageBirdService) send(sms Sms) (SmsResult, error) {
	return SmsResult{success: true}, nil
}
