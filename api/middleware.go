package api

import (
	"log"
	"net/http"
	"time"
)

func Logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		h.ServeHTTP(w, r)
		log.Printf("%s %s %v\n", r.Method, r.URL.Path, time.Since(t))
	})
}

func JSON(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		h.ServeHTTP(w, r)
	})
}
