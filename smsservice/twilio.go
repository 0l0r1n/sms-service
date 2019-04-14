package smsservice

// TwilioService is a gateway for sending messages using twilio
type TwilioService struct {
	next SmsService
}

func (twilio TwilioService) send(sms Sms) (SmsResult, error) {
	return SmsResult{success: true}, nil
}
