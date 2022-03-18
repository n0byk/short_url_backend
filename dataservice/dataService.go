package dataservice

import "github.com/n0byk/short_url_backend/dataservice/entities"

type Repository interface {
	AddURL(key, url string) error
	GetURL(key string) (string, error)
	SetUserData(key, url, user string) error
	GetUserData(user string) ([]entities.URLCatalog, error)
	DBPing() error
}
