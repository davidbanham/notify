package main

import (
	"encoding/json"
	"fmt"
	"github.com/davidbanham/notify/config"
	"github.com/davidbanham/notify/email"
	"github.com/davidbanham/notify/sms"
	"github.com/davidbanham/notify/types"
	"log"
	"net/http"
	"time"
)

func main() {
	topRouter := http.NewServeMux()
	topRouter.HandleFunc("/v1/sms", smsHandler)
	topRouter.HandleFunc("/v1/emails", emailHandler)
	topRouter.HandleFunc("/v1/health", healthHandler)
	topRouter.HandleFunc("/health", healthHandler)

	handler := recoverWrap(topRouter)
	handler = authWrap(handler)

	srv := &http.Server{
		Handler:      handler,
		Addr:         ":" + config.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func smsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var t types.SMS

	err := decoder.Decode(&t)
	if err != nil {
		errRes(w, http.StatusBadRequest, "Error decoding json", err)
		return
	}

	result, err := sms.Send(t)
	if err != nil {
		fmt.Println("Error sending SMS", err)
		errRes(w, http.StatusInternalServerError, "Error sending SMS", err)
		return
	}

	json.NewEncoder(w).Encode(result)
	return
}

func emailHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var t types.Email

	err := decoder.Decode(&t)
	if err != nil {
		errRes(w, http.StatusBadRequest, "Error decoding json", err)
		return
	}

	result, err := email.Send(t)
	if err != nil {
		errRes(w, http.StatusInternalServerError, "Error sending email", err)
		return
	}

	json.NewEncoder(w).Encode(result)
	return
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	successRes(w)
	return
}

func errRes(w http.ResponseWriter, code int, message string, err error) {
	fmt.Println("Error", code, message, err)
	resData := res{
		Success: false,
		Message: message,
		Error:   err,
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
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   error  `json:"error"`
}
