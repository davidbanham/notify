package main

import (
	"fmt"
	"github.com/davidbanham/notify/config"
	"net/http"
)

func authWrap(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if config.DisableAuth {
			h.ServeHTTP(w, r)
			return
		}

		id, pass, ok := r.BasicAuth()

		if !ok {
			w.Header().Set("WWW-Authenticate", `Basic realm="Ordermentum Identity Service (Tapu)"`)
			w.WriteHeader(401)
			return
		}

		if pass != config.AuthSecret {
			fmt.Println("Rejected invalid auth secret", id, pass)
			w.WriteHeader(403)
			return
		}

		h.ServeHTTP(w, r)
		return
	})
}
