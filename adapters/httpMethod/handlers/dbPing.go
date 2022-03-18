package handlers

import (
	"log"
	"net/http"

	config "github.com/n0byk/short_url_backend/config"
)

func DBPing(w http.ResponseWriter, r *http.Request) {
	err := config.AppService.Storage.DBPing()
	if err != nil {
		log.Println("Can't DBPing")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
