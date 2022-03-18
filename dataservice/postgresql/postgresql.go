package postgresql

import (
	"context"

	"github.com/jackc/pgx/v4"
	dataservice "github.com/n0byk/short_url_backend/dataservice"
	entities "github.com/n0byk/short_url_backend/dataservice/entities"
)

type dbRepository struct {
	db *pgx.Conn
}

func (db *dbRepository) AddURL(key, url string) error {
	//some implementation
	return nil
}

func (f *dbRepository) SetUserData(key, url, user string) error {

	//some implementation
	return nil
}

func (f *dbRepository) DBPing() error {
	return f.db.Ping(context.Background())
}

func (m *dbRepository) GetUserData(user string) ([]entities.URLCatalog, error) {

	return []entities.URLCatalog{}, nil
}

func (db *dbRepository) GetURL(key string) (string, error) {
	//some implementation
	return "", nil
}

func NewDBRepository(db *pgx.Conn) dataservice.Repository {
	return &dbRepository{
		db: db,
	}
}
