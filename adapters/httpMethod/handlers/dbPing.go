package handlers

import (
	"fmt"
	"net/http"

	config "github.com/n0byk/short_url_backend/config"
)

func DBPing(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DBPing")

	err := config.AppService.Storage.DBPing()
	if err != nil {
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
}
