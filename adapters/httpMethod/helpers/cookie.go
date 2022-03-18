package helpers

import (
	"net/http"
)

func Cookie(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		SetCookie(w, r)
		next.ServeHTTP(w, r)
	})
}
