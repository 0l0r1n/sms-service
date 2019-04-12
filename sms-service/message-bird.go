package smsservice

// MessageBirdService is a gateway for sending messages using twilio
type MessageBirdService struct {
	next SmsService
}

func (s *MessageBirdService) send(sms Sms) {

}
