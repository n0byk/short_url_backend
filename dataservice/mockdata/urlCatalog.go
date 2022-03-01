package mockdata

import (
	"errors"

	entities "github.com/n0byk/short_url_backend/dataservice/entities"
)

var urlCatalogDB = make(map[string]string)

type URLCatalog entities.URLCatalog

func (f *URLCatalog) AddElement(props entities.URLCatalog) (string, error) {
	urlCatalogDB[props.ShortURL] = props.FullURL

	return props.ShortURL, nil
}

func (f *URLCatalog) GetElement(shortUrl string) (string, error) {
	fullURL, exists := urlCatalogDB[shortUrl]
	if !exists {
		return "", errors.New("Cant_get_URL")
	}

	return fullURL, nil
}
