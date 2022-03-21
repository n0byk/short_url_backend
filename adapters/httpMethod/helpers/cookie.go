package helpers

import (
	"net/http"

	helpers "github.com/n0byk/short_url_backend/helpers"
)

func Cookie(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		http.SetCookie(w, &http.Cookie{Name: "user_id", Value: helpers.GenerateToken(8), Path: "/", Secure: true})

		next.ServeHTTP(w, r)
	})
}
