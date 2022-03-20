package helpers

import (
	"net/http"
	"net/url"

	helpers "github.com/n0byk/short_url_backend/helpers"
)

func JSONResponse(w http.ResponseWriter, responce []byte, httpStatus int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	w.Write(responce)
}

func ValidateURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func SetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("user_id")
	if len(cookie.Value) == 0 {

		ck := http.Cookie{Name: "user_id", Value: helpers.GenerateToken(8), HttpOnly: false}

		http.SetCookie(w, &ck)
	}

}
