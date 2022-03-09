package dataservice

import (
	entities "github.com/n0byk/short_url_backend/dataservice/entities"
)

type URLCatalog interface {
	AddElement(props entities.URLCatalog) (string, error)
	GetElement(props entities.URLCatalog) (string, error)
}
