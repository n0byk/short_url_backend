package dataservice

import "github.com/n0byk/short_url_backend/dataservice/entities"

type Repository interface {
	AddURL(url, user string) (string, error)
	GetURL(key string) (string, error)
	SetUserData(key, url, user string) error
	GetUserData(user string) ([]entities.URLCatalog, error)
	DBPing() error
}
