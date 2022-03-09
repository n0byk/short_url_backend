package main

import (
	"log"
	"net/http"

	httpMethod "github.com/n0byk/short_url_backend/adapters/httpMethod"
	config "github.com/n0byk/short_url_backend/config"
)

func main() {

	log.Print("Started at " + config.AppEnv().ServerAddress)
	log.Fatal(http.ListenAndServe(config.AppEnv().ServerAddress, httpMethod.Routers()))
}
