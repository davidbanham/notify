package main

import (
	"errors"
	"fmt"
	"github.com/davidbanham/notify/config"
	"net/http"
)

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
