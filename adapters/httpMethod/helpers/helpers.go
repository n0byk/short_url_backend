package helpers

import (
	"net/http"
	"net/url"
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
