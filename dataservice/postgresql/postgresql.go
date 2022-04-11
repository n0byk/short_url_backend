package postgresql

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/n0byk/short_url_backend/config"
	dataservice "github.com/n0byk/short_url_backend/dataservice"
	entities "github.com/n0byk/short_url_backend/dataservice/entities"
	"github.com/n0byk/short_url_backend/helpers"
)

type dbRepository struct {
	db *pgx.Conn
}

func (db *dbRepository) AddURL(url, user string) (string, bool, error) {

	key := helpers.GenerateToken(config.AppService.ShortLinkLen)

	_, err := db.db.Exec(context.Background(), `INSERT INTO url_catalog (user_id, full_url, short_url) VALUES($1,$2,$3);`, user, url, key)
	if err != nil {

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				var token string
				err := db.db.QueryRow(context.Background(), "select short_url from url_catalog where full_url=$1 and delete_time is null; ", url).Scan(&token)
				switch err {
				case nil:
					return token, true, nil
				case pgx.ErrNoRows:
					return "", false, nil
				default:
					log.Println(err)
					return "", false, errors.New("DB error")
				}
			}
		} else {
			log.Println(err)
			return "", false, errors.New("DB error")
		}

	}
	return key, false, nil
}

func (db *dbRepository) SetUserData(key, url, user string) error {
	return nil
}

func (db *dbRepository) DBPing() error {
	return db.db.Ping(context.Background())
}

func (db *dbRepository) GetUserData(user string) ([]entities.URLCatalog, error) {

	rows, _ := db.db.Query(context.Background(), "select full_url, short_url from url_catalog where user_id=$1 and delete_time is null; ", user)

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

	err := db.db.QueryRow(context.Background(), "select full_url from url_catalog where short_url=$1 and delete_time is null;", key).Scan(&url)
	switch err {
	case nil:
		return url, nil
	case pgx.ErrNoRows:
		return "", err
	default:
		log.Println(err)
		return "", errors.New("DB error")
	}
}

func (db *dbRepository) BulkDelete(urls []string, userID string, wg *sync.WaitGroup) error {

	test, err := db.db.Exec(context.Background(), "UPDATE url_catalog SET delete_time = now() WHERE short_url IN($1) and user_id = $2;", strings.Join(urls, ", "), userID)
	fmt.Println(test)
	if err != nil {
		log.Println(err)
		return err
	}

	wg.Done()

	return nil
}

func NewDBRepository(db *pgx.Conn) dataservice.Repository {
	return &dbRepository{
		db: db,
	}
}
