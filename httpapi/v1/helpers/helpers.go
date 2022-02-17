package helpers

import (
	"crypto/rand"
	"encoding/base64"
	"net/url"
)

func ValidateUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func GenerateToken(len int) string {
	buff := make([]byte, len)
	rand.Read(buff)
	return base64.RawURLEncoding.EncodeToString(buff)[:len]
}
