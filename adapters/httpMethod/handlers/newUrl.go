package handlers

import (
	"io/ioutil"
	"log"
	"net/http"

	httpMethodHelpers "github.com/n0byk/short_url_backend/adapters/httpMethod/helpers"
	config "github.com/n0byk/short_url_backend/config"
	entities "github.com/n0byk/short_url_backend/dataservice/entities"
	filestorage "github.com/n0byk/short_url_backend/dataservice/filestorage"
	memory "github.com/n0byk/short_url_backend/dataservice/memory"
	helpers "github.com/n0byk/short_url_backend/helpers"
)

func NewURL(w http.ResponseWriter, r *http.Request) {

	urlBytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println("ERR NewUrl - " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !httpMethodHelpers.ValidateURL(string(urlBytes)) {
		log.Println("Validate error - " + string(urlBytes))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token := helpers.GenerateToken(7)

	var adapter memory.URLCatalog

	if len(config.AppEnv().FileStoragePath) > 1 {
		storage, err := filestorage.NewStorageSet(config.AppEnv().FileStoragePath)
		if err != nil {
			log.Printf("NewStorageGet: %v", err)
			http.Error(w, "File Storage Error", http.StatusInternalServerError)
			return
		}
		defer storage.CloseURLCatalog()
		storage.WriteURL(&entities.URLCatalog{ShortURL: token, FullURL: string(urlBytes)})
	} else {
		adapter.AddElement(entities.URLCatalog{ShortURL: token, FullURL: string(urlBytes)})
	}

	w.Header().Set("Content-Type", "application/text")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(config.AppEnv().BaseURL + "/" + token))

}
