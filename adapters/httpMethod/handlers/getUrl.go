package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	config "github.com/n0byk/short_url_backend/config"
	filestorage "github.com/n0byk/short_url_backend/dataservice/filestorage"
	memory "github.com/n0byk/short_url_backend/dataservice/memory"
)

func GetURL(w http.ResponseWriter, r *http.Request) {
	shortURL := mux.Vars(r)["url"]

	var fullURL string
	if len(config.AppEnv().FileStoragePath) > 1 {
		storage, err := filestorage.NewStorageGet(config.AppEnv().FileStoragePath)
		if err != nil {
			log.Fatal(err)
		}
		defer storage.CloseURLCatalog()
		fullURL, err = storage.GetURL(shortURL)
		if err == nil {
			http.Redirect(w, r, fullURL, http.StatusTemporaryRedirect)
			return
		}
		log.Print("Err ", err)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		var adapter memory.URLCatalog
		var err error
		fullURL, err = adapter.GetElement(shortURL)
		if err == nil {
			http.Redirect(w, r, fullURL, http.StatusTemporaryRedirect)
			return
		}
		log.Print("Err ", err)
		w.WriteHeader(http.StatusBadRequest)
	}

}
