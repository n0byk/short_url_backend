package config

import (
	"errors"

	"flag"

	"github.com/caarlos0/env"
)

type appConfig struct {
	ServerAddress   string `env:"SERVER_ADDRESS"  envDefault:"localhost:8080"`
	BaseURL         string `env:"BASE_URL" envDefault:"http://localhost:8080"`
	FileStoragePath string `env:"FILE_STORAGE_PATH"`
}

var appEnv appConfig

func init() {
	err := env.Parse(&appEnv)
	if err != nil {
		panic(err)
	}

	flag.StringVar(&appEnv.ServerAddress, "a", appEnv.ServerAddress, "SERVER_ADDRESS")
	flag.StringVar(&appEnv.BaseURL, "b", appEnv.BaseURL, "BASE_URL")
	flag.Func("f", "FILE_STORAGE_PATH", func(storagePath string) error {
		if storagePath != "" {
			appEnv.FileStoragePath = storagePath
			return nil
		}
		appEnv.FileStoragePath = "url_catalog.db"
		return errors.New("Can't_get_FILE_STORAGE_PATH")
	})
}

func AppEnv() appConfig {
	flag.Parse()
	return appEnv
}
