package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	UrlCatalog "github.com/n0byk/short_url_backend/httpapi/v1/storage/repositories"
)

func GetUrl(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	shortUrl := params["id"]

	fullUrl, exists := UrlCatalog.GetElement(shortUrl)
	if exists {
		http.Redirect(w, r, fullUrl, http.StatusTemporaryRedirect)
		return
	}
	w.WriteHeader(http.StatusBadRequest)

}
