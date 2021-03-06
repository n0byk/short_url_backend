package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	config "github.com/n0byk/short_url_backend/config"
)

//Get url info handler
func GetURL(w http.ResponseWriter, r *http.Request) {

	shortURL := mux.Vars(r)["url"]

	fullURL, err := config.AppService.Storage.GetURL(context.Background(), shortURL)

	if err == nil {
		http.Redirect(w, r, fullURL, http.StatusTemporaryRedirect)
		return
	}
	log.Println("Err ", err)
	w.WriteHeader(http.StatusGone)
}
