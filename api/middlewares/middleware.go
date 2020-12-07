package middlewares

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/triano4/uploaddrive/api/auth"
	"github.com/triano4/uploaddrive/api/responses"
)

//SetMiddlewareJSON function
func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// return "OKOK"
		json.NewEncoder(w).Encode("OKOK")
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
