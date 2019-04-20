package smsservice

// SimpleNotificationService is a gateway for sending messages using AWS
type SimpleNotificationService struct {
}

func (sns *SimpleNotificationService) send(sms SmsMessage) (SmsResult, error) {
	return SmsResult{success: true}, nil
}
