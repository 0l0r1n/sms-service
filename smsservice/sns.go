package smsservice

// SimpleNotificationService is a gateway for sending messages using AWS
type SimpleNotificationService struct {
	next SmsService
}

func (sns *SimpleNotificationService) send(sms Sms) (SmsResult, error) {
	return SmsResult{success: true}, nil
}
