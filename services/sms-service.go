package main

// SmsService represents a Sms notification gateway
type SmsService interface {
	send(sms Sms)
}
