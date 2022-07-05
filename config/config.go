package config

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	dataservice "github.com/n0byk/short_url_backend/dataservice"
)

// App config struct
type AppConfig struct {
	ServerAddress   string `env:"SERVER_ADDRESS"  envDefault:"localhost:8080"`
	BaseURL         string `env:"BASE_URL" envDefault:"http://localhost:8080"`
	FileStoragePath string `env:"FILE_STORAGE_PATH"`
	// DB              string `env:"DATABASE_DSN" envDefault:"postgres://developer:developer@localhost:5432/app?sslmode=disable"`
	DB     string `env:"DATABASE_DSN"`
	TLS    bool   `env:"ENABLE_HTTPS"`
	CONFIG string `env:"CONFIG"`
	TLScrt string `env:"TLScrt"`
	TLSkey string `env:"TLSkey"`
}

type JSONConfig struct {
	FileStoragePath string `json:"file_storage_path"`
	ServerAddress   string `json:"server_address"`
	BaseURL         string `json:"base_url"`
	EnableHTTPS     bool   `json:"enable_https"`
	DB              string `json:"database_dsn"`
}

type Service struct {
	ShortLinkLen int
	BaseURL      string
	Storage      dataservice.Repository
}

var AppService Service

func ShowBuildInfo(buildVersion, buildDate, buildCommit string) {
	if buildVersion == "" {
		buildVersion = "N/A"
	}

	if buildDate == "" {
		buildDate = "N/A"
	}

	if buildCommit == "" {
		buildCommit = "N/A"
	}
	fmt.Printf("Build version: %s\nBuild date: %s\nBuild commit: %s\n", buildVersion, buildDate, buildCommit)

}

func ReadJSONConfig(appEnv AppConfig) *AppConfig {
	if len(appEnv.CONFIG) < 1 {
		return &appEnv
	}

	jsonConfigFile, err := os.OpenFile(appEnv.CONFIG, os.O_RDONLY, 0644)
	conf := JSONConfig{}

	if err != nil {
		log.Println("error reading json config file " + err.Error())
	} else {
		jsonBody, err := io.ReadAll(jsonConfigFile)

		if err != nil {
			log.Println("error reading json config file " + err.Error())
		} else if err = json.Unmarshal(jsonBody, &conf); err != nil {
			log.Println("error reading json config file " + err.Error())
		}
	}

	if appEnv.TLS == false {
		appEnv.TLS = conf.EnableHTTPS
	}

	if len(appEnv.FileStoragePath) < 1 && len(appEnv.FileStoragePath) > 0 {
		appEnv.FileStoragePath = conf.FileStoragePath
	}

	if len(appEnv.ServerAddress) < 1 && len(appEnv.ServerAddress) > 0 {
		appEnv.ServerAddress = conf.ServerAddress
	}

	if len(appEnv.BaseURL) < 1 && len(appEnv.BaseURL) > 0 {
		appEnv.BaseURL = conf.BaseURL
	}

	if len(appEnv.DB) < 1 && len(appEnv.DB) > 0 {
		appEnv.DB = conf.DB
	}

	return &appEnv
}
