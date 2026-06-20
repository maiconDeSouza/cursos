package middleware

import (
	"log"
	"net/http"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s", r.Method)
		next.ServeHTTP(w, r)
	})
}
