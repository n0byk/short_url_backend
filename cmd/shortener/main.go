package main

import (
	"log"
	"net/http"

	httpApi "github.com/n0byk/short_url_backend/httpapi/v1"
)

func main() {

	log.Print("Started at port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", httpApi.V1()))
}
