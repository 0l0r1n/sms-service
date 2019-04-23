package smsservice

// SimpleNotificationService is a gateway for sending messages using AWS
type SimpleNotificationService struct {
}

func (sns *SimpleNotificationService) Send(sms SmsMessage) (SmsResult, error) {
	return SmsResult{success: true}, nil
}
