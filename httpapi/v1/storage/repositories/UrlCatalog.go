package repositories

import (
	"log"

	Storage "github.com/n0byk/short_url_backend/httpapi/v1/storage"
	Entities "github.com/n0byk/short_url_backend/httpapi/v1/storage/entities"
)

type UrlCatalog interface {
	AddElement()
	GetElement()
}

func AddElement(params Entities.AddElement) {
	Storage.UrlCatalogDb[params.ShortUrl] = string(params.UrlBytes)
}

func GetElement(param string) (string, bool) {
	log.Print(Storage.UrlCatalogDb)
	fullUrl, exists := Storage.UrlCatalogDb[param]

	return fullUrl, exists
}
