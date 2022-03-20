package helpers

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

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
	cookie, err := r.Cookie("user_id")

	if err != nil {

		fmt.Println(err)
	}

	if len(cookie.Value) == 0 {
		fmt.Println("new cookie")
		ck := http.Cookie{
			Name:    "user_id",
			Path:    "/",
			Expires: time.Now().AddDate(1, 0, 0), //1 год
			Value:   helpers.GenerateToken(8),
		}

		http.SetCookie(w, &ck)
	}

}
