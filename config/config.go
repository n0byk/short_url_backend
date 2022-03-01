package config

import (
	"github.com/caarlos0/env"
)

type appConfig struct {
	ServerAddress   string `env:"SERVER_ADDRESS" envDefault:"localhost:8080"`
	BaseURL         string `env:"BASE_URL" envDefault:"http://localhost:8080"`
	FileStoragePath string `env:"FILE_STORAGE_PATH" envDefault:"url_catalog.db"`
}

func AppEnv() appConfig {

	appEnv := appConfig{}
	if err := env.Parse(&appEnv); err != nil {
		panic(err)
	}
	return appEnv
}
