package middlewares

import (
	"log"
	"net/http"

	"api/src/auth"
	"api/src/responses"
)

// Logger represents requests on terminal
func Logger(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s ", r.Method, r.RequestURI, r.Host)
		nextFunction(w, r)
	}
}

// Authenticate verify user execute request auth
func Authenticate(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := auth.ValidateToken(r); err != nil {
			responses.Err(w, http.StatusUnauthorized, err)
			return
		}
		nextFunction(w, r)
	}
}
