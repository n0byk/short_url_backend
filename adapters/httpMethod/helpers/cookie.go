package helpers

import (
	"net/http"

	"github.com/n0byk/short_url_backend/helpers"
)

func Cookie(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		_, err := r.Cookie("user_id")

		if err == http.ErrNoCookie {
			token := helpers.GenerateToken(10)
			r.AddCookie(&http.Cookie{Name: "user_id", Value: token, Path: "/", Secure: false})
			http.SetCookie(w, &http.Cookie{Name: "user_id", Value: token, Path: "/", Secure: false})
		}

		next.ServeHTTP(w, r)
	})
}
