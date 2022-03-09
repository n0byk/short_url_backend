package helpers

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateToken(len int) string {
	buff := make([]byte, len)
	rand.Read(buff)
	return base64.RawURLEncoding.EncodeToString(buff)[:len]
}
