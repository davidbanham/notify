package main

import (
	"encoding/json"
	"github.com/davidbanham/notify/email"
	"github.com/davidbanham/notify/sms"
	"github.com/davidbanham/notify/types"
	"github.com/davidbanham/required_env"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

func init() {
	required_env.Ensure(map[string]string{
		"PORT":                  "",
		"NOTIFY_EMAIL_PROVIDER": "",
		"NOTIFY_SMS_PROVIDER":   "",
	})

	switch os.Getenv("NOTIFY_EMAIL_PROVIDER") {
	case "gmail":
		required_env.Ensure(map[string]string{
			"NOTIFY_EMAIL_SMTP_PASS": "",
			"NOTIFY_EMAIL_FROM":      "",
		})
	}
	switch os.Getenv("NOTIFY_SMS_PROVIDER") {
	case "amazon":
		required_env.Ensure(map[string]string{
			"AWS_ACCESS_KEY_ID":     "",
			"AWS_SECRET_ACCESS_KEY": "",
			"AWS_REGION":            "",
		})
	}
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/sms", SmsHandler).
		Methods("POST")

	r.HandleFunc("/email", EmailHandler).
		Methods("POST")

	srv := &http.Server{
		Handler:      r,
		Addr:         ":" + os.Getenv("PORT"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func SmsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var t types.SMS

	err := decoder.Decode(&t)
	if err != nil {
		errRes(w, http.StatusBadRequest)
		return
	}

	err = sms.Send(t)
	if err != nil {
		errRes(w, http.StatusInternalServerError)
		return
	}

	successRes(w)
	return
}

func EmailHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var t types.Email

	err := decoder.Decode(&t)
	if err != nil {
		errRes(w, http.StatusBadRequest)
		return
	}

	err = email.Send(t)
	if err != nil {
		errRes(w, http.StatusInternalServerError)
		return
	}

	successRes(w)
	return
}

func errRes(w http.ResponseWriter, code int) {
	resData := res{
		Success: false,
	}

	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resData)
}

func successRes(w http.ResponseWriter) {
	resData := res{
		Success: true,
	}

	json.NewEncoder(w).Encode(resData)
}

type res struct {
	Success bool `json:"success"`
}
