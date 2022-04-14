package filestorage

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/n0byk/short_url_backend/config"
	dataservice "github.com/n0byk/short_url_backend/dataservice"
	entities "github.com/n0byk/short_url_backend/dataservice/entities"
	"github.com/n0byk/short_url_backend/helpers"
)

type fileRepository struct {
	f        *os.File
	urlsDB   map[string]entities.URLdb
	userData map[string][]entities.URLCatalog
}

func (f *fileRepository) AddURL(ctx context.Context, url, user string) (string, bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	key := helpers.GenerateToken(config.AppService.ShortLinkLen)
	if len(f.urlsDB) == 0 {
		byteValue, _ := ioutil.ReadAll(f.f)
		json.Unmarshal([]byte(byteValue), &f.urlsDB)
	}

	f.urlsDB[key] = entities.URLdb{FullURL: url, UserID: user}

	jsonData, err := json.Marshal(f.urlsDB)

	if err != nil {
		log.Println(err)
		return "", false, errors.New("filestorage: some troubles while adding new url")
	}
	f.f.Truncate(0)
	f.f.Write(jsonData)
	return key, false, nil
}

func (f *fileRepository) DBPing() error {
	return errors.New("filestorage: only for Postgresql")
}

func (f *fileRepository) SetUserData(ctx context.Context, key, url, user string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	f.userData[user] = append(f.userData[user], entities.URLCatalog{ShortURL: key, FullURL: url})

	return nil
}

func (f *fileRepository) GetUserData(ctx context.Context, user string) ([]entities.URLCatalog, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	data, exists := f.userData[user]
	if !exists {
		return []entities.URLCatalog{}, errors.New("Cant_get_user_info")
	}

	return data, nil
}

func (f *fileRepository) GetURL(ctx context.Context, key string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	var urls = f.urlsDB
	byteValue, _ := ioutil.ReadAll(f.f)
	json.Unmarshal([]byte(byteValue), &urls)

	fullURL, exists := urls[key]
	if !exists {
		return "", errors.New("Cant_get_URL")
	}

	return fullURL.FullURL, nil
}

func (f *fileRepository) BulkDelete(ctx context.Context, urls []string, userID string) error {
	for _, v := range urls {
		delete(f.urlsDB, v)

	}
	return nil
}

func NewFileRepository(f *os.File) dataservice.Repository {
	return &fileRepository{
		f:        f,
		urlsDB:   make(map[string]entities.URLdb),
		userData: make(map[string][]entities.URLCatalog),
	}
}
