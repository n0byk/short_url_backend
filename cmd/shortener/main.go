package main

import (
	"log"
	"net/http"

	version1 "github.com/n0byk/short_url_backend/adapter/http/version1"
)

func main() {

	log.Print("Started at port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", version1.V1()))
}
