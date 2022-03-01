package config

import (
	"github.com/caarlos0/env"
)

type appConfig struct {
	SERVER_ADDRESS string `env:"SERVER_ADDRESS" envDefault:"localhost:8080"`
	BASE_URL       string `env:"BASE_URL" envDefault:"http://localhost:8080/"`
}

func AppEnv() appConfig {

	appEnv := appConfig{}
	if err := env.Parse(&appEnv); err != nil {
		panic(err)
	}
	return appEnv
}
