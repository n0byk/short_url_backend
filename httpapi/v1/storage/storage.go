package storage

import (
	"encoding/base64"
	"math"
	"math/rand"
)

var UrlCatalogDb = make(map[string]string)

func RandomToken(l int) string {
	buff := make([]byte, int(math.Ceil(float64(l)/float64(1.3333123441433))))
	rand.Read(buff)
	str := base64.RawURLEncoding.EncodeToString(buff)
	return str[:l]
}
