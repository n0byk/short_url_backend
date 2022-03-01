package helpers

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"net/url"
)

func ValidateURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func GenerateToken(len int) string {
	buff := make([]byte, len)
	rand.Read(buff)
	return base64.RawURLEncoding.EncodeToString(buff)[:len]
}

func JSONResponse(w http.ResponseWriter, responce []byte, httpStatus int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	w.Write(responce)
}
