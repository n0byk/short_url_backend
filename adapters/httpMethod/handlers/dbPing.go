package handlers

import (
	"net/http"

	config "github.com/n0byk/short_url_backend/config"
)

func DBPing(w http.ResponseWriter, r *http.Request) {
	err := config.AppService.Storage.DBPing()
	if err != nil {
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
}
