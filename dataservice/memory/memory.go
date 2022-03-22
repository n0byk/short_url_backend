package memory

import (
	"errors"

	"github.com/n0byk/short_url_backend/config"
	dataservice "github.com/n0byk/short_url_backend/dataservice"
	entities "github.com/n0byk/short_url_backend/dataservice/entities"
	"github.com/n0byk/short_url_backend/helpers"
)

type memoryRepository struct {
	urlsDB   map[string]string
	userData map[string][]entities.URLCatalog
}

func (m *memoryRepository) AddURL(url, user string) (string, error) {
	key := helpers.GenerateToken(config.AppService.ShortLinkLen)
	m.urlsDB[key] = url
	return key, nil
}

func (m *memoryRepository) DBPing() error {
	return errors.New("only for Postgresql")
}

func (m *memoryRepository) SetUserData(key, url, user string) error {

	m.userData[user] = append(m.userData[user], entities.URLCatalog{ShortURL: key, FullURL: url})

	return nil
}

func (m *memoryRepository) GetUserData(user string) ([]entities.URLCatalog, error) {

	data, exists := m.userData[user]
	if !exists {
		return []entities.URLCatalog{}, errors.New("Cant_get_user_info")
	}

	return data, nil
}

func (m *memoryRepository) GetURL(key string) (string, error) {
	fullURL, exists := m.urlsDB[key]
	if !exists {
		return "", errors.New("Cant_get_URL")
	}

	return fullURL, nil
}

func NewMemoryRepository() dataservice.Repository {
	return &memoryRepository{urlsDB: make(map[string]string), userData: make(map[string][]entities.URLCatalog)}
}
