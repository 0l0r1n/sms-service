package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/0l0r1n/sms-service/smsservice"
)

// SmsHandler accepts a request to send sms and dispatches it to the chain of responsibility of sms services
func SmsHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var sms smsservice.SmsMessage
	err = json.Unmarshal(b, &sms)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	result := smsservice.GetChainOfServices().Send(&sms)
	if !result.Success {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusAccepted)
	}

	output, _ := json.Marshal(result)

	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
