package filestorage

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"

	"github.com/n0byk/short_url_backend/config"
	dataservice "github.com/n0byk/short_url_backend/dataservice"
	entities "github.com/n0byk/short_url_backend/dataservice/entities"
	"github.com/n0byk/short_url_backend/helpers"
)

type fileRepository struct {
	f        *os.File
	urlsDB   map[string]string
	userData map[string][]entities.URLCatalog
}

func (f *fileRepository) AddURL(url, user string) (string, bool, error) {
	key := helpers.GenerateToken(config.AppService.ShortLinkLen)
	if len(f.urlsDB) == 0 {
		byteValue, _ := ioutil.ReadAll(f.f)
		json.Unmarshal([]byte(byteValue), &f.urlsDB)
	}

	f.urlsDB[key] = url

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

func (f *fileRepository) SetUserData(key, url, user string) error {

	f.userData[user] = append(f.userData[user], entities.URLCatalog{ShortURL: key, FullURL: url})

	return nil
}

func (f *fileRepository) GetUserData(user string) ([]entities.URLCatalog, error) {

	data, exists := f.userData[user]
	if !exists {
		return []entities.URLCatalog{}, errors.New("Cant_get_user_info")
	}

	return data, nil
}

func (f *fileRepository) GetURL(key string) (string, error) {
	var urls = f.urlsDB
	byteValue, _ := ioutil.ReadAll(f.f)
	json.Unmarshal([]byte(byteValue), &urls)

	fullURL, exists := urls[key]
	if !exists {
		return "", errors.New("Cant_get_URL")
	}

	return fullURL, nil
}

func NewFileRepository(f *os.File) dataservice.Repository {
	return &fileRepository{
		f:        f,
		urlsDB:   make(map[string]string),
		userData: make(map[string][]entities.URLCatalog),
	}
}
