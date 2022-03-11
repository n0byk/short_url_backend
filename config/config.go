package config

import (
	dataservice "github.com/n0byk/short_url_backend/dataservice"
)

type AppConfig struct {
	ServerAddress   string `env:"SERVER_ADDRESS"  envDefault:"localhost:8080"`
	BaseURL         string `env:"BASE_URL" envDefault:"http://localhost:8080"`
	FileStoragePath string `env:"FILE_STORAGE_PATH"`
	DB              string `env:"DATABASE_CONNECTION_STRING"`
}

type Service struct {
	ShortLinkLen int
	BaseURL      string
	Storage      dataservice.Repository
}

var AppService Service
