package middlewares

import (
	"errors"
	"net/http"

	"github.com/triano4/uploaddrive/api/auth"
	"github.com/triano4/uploaddrive/api/responses"
)

//SetMiddlewareJSON function
func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token")
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "OPTIONS" {
			w.Write([]byte("allowed"))
			return
		}

		next(w, r)
	}
}

//SetMiddlewareAuthentication function
func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		next(w, r)
	}
}
