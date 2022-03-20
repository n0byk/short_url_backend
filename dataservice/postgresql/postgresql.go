package postgresql

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
	dataservice "github.com/n0byk/short_url_backend/dataservice"
	entities "github.com/n0byk/short_url_backend/dataservice/entities"
)

type dbRepository struct {
	db *pgx.Conn
}

func (db *dbRepository) AddURL(key, url, user string) error {
	_, err := db.db.Exec(context.Background(), `INSERT INTO url_catalog (user_id, full_url, short_url) VALUES($1,$2,$3);`, user, url, key)
	if err != nil {
		log.Println(err)
		return errors.New("DB error")
	}
	return nil
}

func (db *dbRepository) SetUserData(key, url, user string) error {
	return nil
}

func (db *dbRepository) DBPing() error {
	return db.db.Ping(context.Background())
}

func (db *dbRepository) GetUserData(user string) ([]entities.URLCatalog, error) {

	rows, _ := db.db.Query(context.Background(), "select full_url, short_url from url_catalog where user_id=$1; ", user)

	var urlCatalog []entities.URLCatalog

	for rows.Next() {
		var r entities.URLCatalog
		err := rows.Scan(&r.FullURL, &r.ShortURL)
		if err != nil {
			log.Fatal(err)
		}
		urlCatalog = append(urlCatalog, r)
	}
	if err := rows.Err(); err != nil {
		return []entities.URLCatalog{}, errors.New("DB error")
	}

	return urlCatalog, nil

}

func (db *dbRepository) GetURL(key string) (string, error) {
	var url string
	err := db.db.QueryRow(context.Background(), "select full_url from url_catalog where short_url=$1; ", key).Scan(&url)
	switch err {
	case nil:
		fmt.Println(url)
		return url, nil
	case pgx.ErrNoRows:
		return "", nil
	default:
		log.Println(err)
		return "", errors.New("DB error")
	}
}

func NewDBRepository(db *pgx.Conn) dataservice.Repository {
	return &dbRepository{
		db: db,
	}
}
