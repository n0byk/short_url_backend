package postgresql

import (
	"database/sql"

	dataservice "github.com/n0byk/short_url_backend/dataservice"
	entities "github.com/n0byk/short_url_backend/dataservice/entities"
)

type dbRepository struct {
	db *sql.DB
}

func (db *dbRepository) AddURL(key, url string) error {
	//some implementation
	return nil
}

func (f *dbRepository) SetUserData(key, url, user string) error {

	//some implementation
	return nil
}

func (m *dbRepository) GetUserData(user string) ([]entities.URLCatalog, error) {

	return nil, nil
}

func (db *dbRepository) GetURL(key string) (string, error) {
	//some implementation
	return "", nil
}

func NewDBRepository(db *sql.DB) dataservice.Repository {
	return &dbRepository{
		db: db,
	}
}
