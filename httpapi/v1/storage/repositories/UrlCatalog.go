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
	storage.URLCatalogDB[params.ShortURL] = string(params.URLBytes)
}

func GetElement(param string) (string, error) {
	log.Print(storage.URLCatalogDB)

	fullURL, exists := storage.URLCatalogDB[param]
	if !exists {
		return "", errors.New("Cant get URL")
	}

	return fullURL, nil
}
