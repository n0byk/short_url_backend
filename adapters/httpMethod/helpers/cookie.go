package helpers

import (
	"bytes"
	"io"
	"net/http"

	helpers "github.com/n0byk/short_url_backend/helpers"
)

type responsewriter struct {
	w    http.ResponseWriter
	buf  bytes.Buffer
	code int
}

func (rw *responsewriter) Header() http.Header {
	return rw.w.Header()
}

func (rw *responsewriter) WriteHeader(statusCode int) {
	rw.code = statusCode
}

func (rw *responsewriter) Write(data []byte) (int, error) {
	return rw.buf.Write(data)
}

func (rw *responsewriter) Done() (int64, error) {
	if rw.code > 0 {
		rw.w.WriteHeader(rw.code)
	}
	return io.Copy(rw.w, &rw.buf)
}

func Cookie(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := &responsewriter{w: w}

		_, err := r.Cookie("user_id")

		if err != nil {
			http.SetCookie(rw, &http.Cookie{Name: "user_id", Value: helpers.GenerateToken(8), Path: "/", Secure: true})
		}

		next.ServeHTTP(rw, r)
	})
}
