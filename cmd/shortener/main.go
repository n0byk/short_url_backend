package main

import (
	"log"
	"net/http"

	version1 "github.com/n0byk/short_url_backend/adapter/http/version1"
	config "github.com/n0byk/short_url_backend/config"
)

func main() {

	log.Print("Started at port 8080")
	log.Fatal(http.ListenAndServe(config.AppEnv().SERVER_ADDRESS, version1.V1()))
}
