package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	urlCatalog "github.com/n0byk/short_url_backend/httpapi/v1/storage/repositories"
)

func GetURL(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	shortURL := params["id"]

	fullURL, err := urlCatalog.GetElement(shortURL)
	if err == nil {

		http.Redirect(w, r, fullURL, http.StatusTemporaryRedirect)
		return
	}
	log.Print("Err ", err)
	w.WriteHeader(http.StatusBadRequest)

}
