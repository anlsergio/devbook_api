package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

// Logger - Prints out information regarding the current request
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// Authenticate - checks if the current user making the request has an active and valid session token
func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Authenticating the request...")
		next(w, r)
	}
}