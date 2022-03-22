package config

import (
	dataservice "github.com/n0byk/short_url_backend/dataservice"
)

type AppConfig struct {
	ServerAddress   string `env:"SERVER_ADDRESS"  envDefault:"localhost:8080"`
	BaseURL         string `env:"BASE_URL" envDefault:"http://localhost:8080"`
	FileStoragePath string `env:"FILE_STORAGE_PATH"`
	DB              string `env:"DATABASE_DSN" envDefault:"postgres://developer:developer@localhost:5432/app?sslmode=disable"`
	// DB string `env:"DATABASE_DSN"`
}

type Service struct {
	ShortLinkLen int
	BaseURL      string
	Storage      dataservice.Repository
}

var AppService Service
