package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/caarlos0/env"

	httpMethod "github.com/n0byk/short_url_backend/adapters/httpMethod"
	config "github.com/n0byk/short_url_backend/config"
	dataservice "github.com/n0byk/short_url_backend/dataservice"
	filestorage "github.com/n0byk/short_url_backend/dataservice/filestorage"
	memory "github.com/n0byk/short_url_backend/dataservice/memory"
	postgresql "github.com/n0byk/short_url_backend/dataservice/postgresql"
)

var appEnv config.AppConfig

func init() {
	flag.StringVar(&appEnv.ServerAddress, "a", "localhost:8080", "SERVER_ADDRESS")
	flag.StringVar(&appEnv.BaseURL, "b", "http://localhost:8080", "BASE_URL")
	flag.StringVar(&appEnv.FileStoragePath, "f", "", "FILE_STORAGE_PATH")
	flag.StringVar(&appEnv.DB, "d", "", "DATABASE_CONNECTION_STRING")

	if err := env.Parse(&appEnv); err != nil {
		log.Fatalf("Unset vars: %v", err)
	}
}

func main() {
	flag.Parse()
	var storage dataservice.Repository

	if appEnv.FileStoragePath != "" {
		f, err := os.OpenFile(appEnv.FileStoragePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		storage = filestorage.NewFileRepository(f)
	}

	if appEnv.DB != "" {
		db, err := sql.Open("postgres", "")
		if err != nil {
			log.Fatal(err)
		}

		storage = postgresql.NewDBRepository(db)
	}

	if appEnv.DB == "" && appEnv.FileStoragePath == "" {

		storage = memory.NewMemoryRepository()

	}

	config.AppService = config.Service{ShortLinkLen: 7, BaseURL: appEnv.BaseURL, Storage: storage}

	log.Print("Started at " + appEnv.ServerAddress)
	// log.Fatal(http.ListenAndServe(appEnv.ServerAddress, httpMethodhelpers.Gzip(httpMethod.Routers())))
	log.Fatal(http.ListenAndServe(appEnv.ServerAddress, httpMethod.Routers()))

}
