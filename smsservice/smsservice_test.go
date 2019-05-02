package smsservice_test

import (
	"testing"

	"github.com/0l0r1n/sms-service/smsservice"
)

func TestGetChainOfServices(t *testing.T) {
	chain := smsservice.GetChainOfServices()
	expected := smsservice.SmsResult{Success: true}
	actual := chain.Send(nil)
	if expected != actual {
		t.Error()
	}
}
