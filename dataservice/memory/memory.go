package memory

import (
	"errors"

	dataservice "github.com/n0byk/short_url_backend/dataservice"
)

type memoryRepository struct{ db map[string]string }

func (m *memoryRepository) AddURL(key, url string) error {
	m.db[key] = url
	return nil
}

func (m *memoryRepository) GetURL(key string) (string, error) {
	fullURL, exists := m.db[key]
	if !exists {
		return "", errors.New("Cant_get_URL")
	}

	return fullURL, nil
}

func NewMemoryRepository() dataservice.Repository {
	return &memoryRepository{db: make(map[string]string)}
}
