package filestorage

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"

	dataservice "github.com/n0byk/short_url_backend/dataservice"
)

type fileRepository struct {
	f  *os.File
	db map[string]string
}

func (f *fileRepository) AddURL(key, url string) error {
	if len(f.db) == 0 {
		byteValue, _ := ioutil.ReadAll(f.f)
		json.Unmarshal([]byte(byteValue), &f.db)
	}

	f.db[key] = url

	jsonData, err := json.Marshal(f.db)

	if err != nil {
		log.Println(err)
		return errors.New("filestorage: some troubles while adding new url")
	}
	f.f.Truncate(0)
	f.f.Write(jsonData)
	return nil
}

func (f *fileRepository) GetURL(key string) (string, error) {
	var urls = f.db
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
		f:  f,
		db: make(map[string]string),
	}
}
