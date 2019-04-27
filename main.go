package main

import (
	"net/http"

	"github.com/0l0r1n/sms-service/api"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/sms/", api.SmsHandler)
	http.Handle("/", r)

	http.ListenAndServe("localhost:8080", r)
}
