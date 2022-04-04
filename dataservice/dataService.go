package dataservice

import (
	"net/url"

	"github.com/n0byk/short_url_backend/dataservice/entities"
)

type Repository interface {
	AddURL(url, user string) (string, bool, error)
	GetURL(key string) (string, error)
	SetUserData(key, url, user string) error
	GetUserData(user string) ([]entities.URLCatalog, error)
	DBPing() error
	BulkDelete(urls url.Values, userID string) error
}
