package postgresql

import (
	"database/sql"

	dataservice "github.com/n0byk/short_url_backend/dataservice"
)

type dbRepository struct {
	db *sql.DB
}

func (db *dbRepository) AddURL(key, url string) error {
	//some implementation
	return nil
}

func (db *dbRepository) GetURL(key string) (string, error) {
	//some implementation
	return "", nil
}

func NewDbRepository(db *sql.DB) dataservice.Repository {
	return &dbRepository{
		db: db,
	}
}
