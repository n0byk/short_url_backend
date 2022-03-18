package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/caarlos0/env"
	"github.com/jackc/pgx/v4"

	httpMethod "github.com/n0byk/short_url_backend/adapters/httpMethod"
	httpMethodhelpers "github.com/n0byk/short_url_backend/adapters/httpMethod/helpers"
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
	flag.StringVar(&appEnv.DB, "d", "", "DATABASE_DSN")

	if err := env.Parse(&appEnv); err != nil {
		log.Fatalf("Unset vars: %v", err)
	}
}

func main() {
	flag.Parse()
	var storage dataservice.Repository

	if appEnv.FileStoragePath != "" {
		log.Println("File storage init.")
		f, err := os.OpenFile(appEnv.FileStoragePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		storage = filestorage.NewFileRepository(f)
	}

	if appEnv.DB != "" {
		log.Println("Postgres storage init.")
		conn, err := pgx.Connect(context.Background(), appEnv.DB)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			os.Exit(1)
		}
		storage = postgresql.NewDBRepository(conn)
		defer conn.Close(context.Background())
	}

	if appEnv.DB == "" && appEnv.FileStoragePath == "" {
		log.Println("Memory storage init.")
		storage = memory.NewMemoryRepository()

	}

	config.AppService = config.Service{ShortLinkLen: 7, BaseURL: appEnv.BaseURL, Storage: storage}

	log.Print("Started at " + appEnv.ServerAddress)
	log.Fatal(http.ListenAndServe(appEnv.ServerAddress, httpMethodhelpers.Cookie(httpMethodhelpers.Gzip(httpMethod.Routers()))))

}
