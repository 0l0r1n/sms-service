package smsservice

// TwilioService is a gateway for sending messages using twilio
type TwilioService struct {
	next SmsService
}

func (s *TwilioService) send(sms Sms) {

}
