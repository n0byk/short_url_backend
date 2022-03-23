package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	config "github.com/n0byk/short_url_backend/config"
)

func GetURL(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetURL")

	shortURL := mux.Vars(r)["url"]

	fullURL, err := config.AppService.Storage.GetURL(shortURL)
	if err == nil {
		http.Redirect(w, r, fullURL, http.StatusTemporaryRedirect)
		return
	}
	log.Println("Err ", err)
	w.WriteHeader(http.StatusBadRequest)
}
