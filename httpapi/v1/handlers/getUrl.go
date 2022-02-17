package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	urlCatalog "github.com/n0byk/short_url_backend/httpapi/v1/storage/repositories"
)

func GetURL(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	shortUrl := params["id"]

	fullUrl, err := urlCatalog.GetElement(shortUrl)
	if err == nil {

		http.Redirect(w, r, fullUrl, http.StatusTemporaryRedirect)
		return
	}
	log.Print("Cant get URL", err)
	w.WriteHeader(http.StatusBadRequest)

}
