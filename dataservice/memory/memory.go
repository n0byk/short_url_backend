package memory

import (
	"context"
	"errors"
	"time"

	"github.com/n0byk/short_url_backend/config"
	dataservice "github.com/n0byk/short_url_backend/dataservice"
	entities "github.com/n0byk/short_url_backend/dataservice/entities"
	"github.com/n0byk/short_url_backend/helpers"
)

type memoryRepository struct {
	urlsDB   map[string]entities.URLdb
	userData map[string][]entities.URLCatalog
}

func (m *memoryRepository) AddURL(ctx context.Context, url, user string) (string, bool, error) {
	_, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	key := helpers.GenerateToken(config.AppService.ShortLinkLen)
	m.urlsDB[key] = entities.URLdb{FullURL: url, UserID: user}
	return key, false, nil
}

func (m *memoryRepository) DBPing() error {
	return errors.New("only for Postgresql")
}

func (m *memoryRepository) SetUserData(ctx context.Context, key, url, user string) error {
	_, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	m.userData[user] = append(m.userData[user], entities.URLCatalog{ShortURL: url, FullURL: key})

	return nil
}

func (m *memoryRepository) GetUserData(ctx context.Context, user string) ([]entities.URLCatalog, error) {
	_, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	data, exists := m.userData[user]
	if !exists {
		return []entities.URLCatalog{}, errors.New("Cant_get_user_info")
	}

	return data, nil
}

func (m *memoryRepository) GetURL(ctx context.Context, key string) (string, error) {
	_, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	fullURL, exists := m.urlsDB[key]
	if !exists {
		return "", errors.New("Cant_get_URL")
	}

	return fullURL.FullURL, nil
}

func (m *memoryRepository) BulkDelete(ctx context.Context, urls []string, userID string) error {
	_, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	for _, v := range urls {
		delete(m.urlsDB, v)

	}
	return nil
}

func NewMemoryRepository() dataservice.Repository {
	return &memoryRepository{urlsDB: make(map[string]entities.URLdb), userData: make(map[string][]entities.URLCatalog)}
}
