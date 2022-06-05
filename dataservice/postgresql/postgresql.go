package postgresql

import (
	"context"
	"errors"
	"log"
	"strings"
	"time"

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

func (db *dbRepository) AddURL(ctx context.Context, url, user string) (string, bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	key := helpers.GenerateToken(config.AppService.ShortLinkLen)

	_, err := db.db.Exec(ctx, `INSERT INTO url_catalog (user_id, full_url, short_url) VALUES($1,$2,$3);`, user, url, key)
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

func (db *dbRepository) SetUserData(ctx context.Context, key, url, user string) error {
	return nil
}

func (db *dbRepository) DBPing() error {
	return db.db.Ping(context.Background())
}

func (db *dbRepository) GetUserData(ctx context.Context, user string) ([]entities.URLCatalog, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	rows, _ := db.db.Query(ctx, "select full_url, short_url from url_catalog where user_id=$1 and delete_time is null; ", user)

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

func (db *dbRepository) GetURL(ctx context.Context, key string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var url string

	err := db.db.QueryRow(ctx, "select full_url from url_catalog where short_url=$1 and delete_time is null;", key).Scan(&url)
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

func (db *dbRepository) BulkDelete(ctx context.Context, urls []string, userID string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	_, err := db.db.Query(ctx, "UPDATE url_catalog SET delete_time = now() WHERE short_url IN($1) and user_id = $2;", strings.Join(urls, ","), userID)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func NewDBRepository(db *pgx.Conn) dataservice.Repository {
	return &dbRepository{
		db: db,
	}
}
