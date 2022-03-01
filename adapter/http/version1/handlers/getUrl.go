package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	mockdata "github.com/n0byk/short_url_backend/dataservice/mockdata"
)

func GetURL(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	shortURL := params["url"]

	var adapter mockdata.URLCatalog

	fullURL, err := adapter.GetElement(shortURL)
	if err == nil {

		http.Redirect(w, r, fullURL, http.StatusTemporaryRedirect)
		return
	}
	log.Print("Err ", err)
	w.WriteHeader(http.StatusBadRequest)

}
