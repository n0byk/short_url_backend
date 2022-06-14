package config

import (
	"fmt"

	dataservice "github.com/n0byk/short_url_backend/dataservice"
)

// App config struct
type AppConfig struct {
	ServerAddress   string `env:"SERVER_ADDRESS"  envDefault:"localhost:8080"`
	BaseURL         string `env:"BASE_URL" envDefault:"http://localhost:8080"`
	FileStoragePath string `env:"FILE_STORAGE_PATH"`
	// DB              string `env:"DATABASE_DSN" envDefault:"postgres://developer:developer@localhost:5432/app?sslmode=disable"`
	DB string `env:"DATABASE_DSN"`
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
