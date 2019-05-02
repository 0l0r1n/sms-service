package smsservice

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var (
	accountSid   = os.Getenv("TWILIO_ACCOUNT_SID")
	authToken    = os.Getenv("TWILIO_AUTH_TOKEN")
	twilioNumber = os.Getenv("TWILIO_NUMBER")
	urlStr       = "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"
)

type twilioService struct {
	next   SmsService
	client *http.Client
}

func (twilio *twilioService) Send(message *SmsMessage) SmsResult {
	resp, _ := twilio.client.Do(buildRequest(message.Recipient, twilioNumber, message.Body))
	var data map[string]interface{}
	decoder := json.NewDecoder(resp.Body)
	err := decoder.Decode(&data)
	if resp.StatusCode >= 400 && twilio.next == nil {
		return errorResult()
	}
	if err == nil {
		return SmsResult{Success: true}
	}
	return twilio.next.Send(message)

}

func newTwilio(next SmsService) SmsService {
	return &twilioService{
		next:   next,
		client: &http.Client{},
	}
}

func buildRequest(to, from, body string) *http.Request {
	msgData := url.Values{}
	msgData.Set("To", to)
	msgData.Set("From", from)
	msgData.Set("Body", body)
	msgDataReader := *strings.NewReader(msgData.Encode())
	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return req
}
