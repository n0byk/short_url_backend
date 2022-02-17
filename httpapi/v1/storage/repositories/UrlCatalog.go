package repositories

import (
	"errors"
	"log"

	storage "github.com/n0byk/short_url_backend/httpapi/v1/storage"
	entities "github.com/n0byk/short_url_backend/httpapi/v1/storage/entities"
)

type UrlCatalog interface {
	AddElement()
	GetElement() (string, error)
}

func AddElement(params entities.AddElement) {
	storage.UrlCatalogDb[params.ShortUrl] = string(params.UrlBytes)
}

func GetElement(param string) (string, error) {
	log.Print(storage.UrlCatalogDb)
	fullUrl, exists := storage.UrlCatalogDb[param]
	log.Print(exists)

	if !exists {
		return "", errors.New("Cant get URL")
	}

	return fullUrl, nil
}
