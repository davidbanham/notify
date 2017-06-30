package main

import (
	"encoding/json"
	"errors"
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
	topRouter.HandleFunc("/sms", smsHandler)
	topRouter.HandleFunc("/email", emailHandler)
	topRouter.HandleFunc("/health", healthHandler)

	recoveredHandler := recoverWrap(topRouter)
	srv := &http.Server{
		Handler:      recoveredHandler,
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

	err = sms.Send(t)
	if err != nil {
		errRes(w, http.StatusInternalServerError, "Error sending SMS", err)
		return
	}

	successRes(w)
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

	err = email.Send(t)
	if err != nil {
		errRes(w, http.StatusInternalServerError, "Error sending email", err)
		return
	}

	successRes(w)
	return
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	successRes(w)
	return
}

func errRes(w http.ResponseWriter, code int, message string, err error) {
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

func recoverWrap(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		defer func() {
			r := recover()
			if r != nil {
				switch t := r.(type) {
				case string:
					err = errors.New(t)
				case error:
					err = t
				default:
					err = errors.New("Unknown error")
				}
				fmt.Println("UNHANDLED PANIC", err)
				errText := ""
				if config.Testing == "true" {
					errText = err.Error()
				}
				errRes(w, 500, errText, err)
			}
		}()
		h.ServeHTTP(w, r)
	})
}
