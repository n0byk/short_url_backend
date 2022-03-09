package config

import (
	"fmt"
	"log"
	"sync"

	"flag"

	"github.com/caarlos0/env"
)

type appConfig struct {
	ServerAddress   string `env:"SERVER_ADDRESS"  envDefault:"localhost:8080"`
	BaseURL         string `env:"BASE_URL" envDefault:"http://localhost:8080"`
	FileStoragePath string `env:"FILE_STORAGE_PATH"`
}

var once sync.Once
var appEnv *appConfig

func AppEnv() *appConfig {
	once.Do(func() {
		appEnv = &appConfig{}
		flag.StringVar(&appEnv.ServerAddress, "a", "localhost:8080", "SERVER_ADDRESS")
		flag.StringVar(&appEnv.BaseURL, "b", "http://localhost:8080", "BASE_URL")
		flag.StringVar(&appEnv.FileStoragePath, "f", "url_catalog.db", "FILE_STORAGE_PATH")
		flag.Parse()

		if err := env.Parse(appEnv); err != nil {
			log.Fatalf("Unset vars: %v", err)
		}
		fmt.Print(appEnv)
	})
	return appEnv

}
