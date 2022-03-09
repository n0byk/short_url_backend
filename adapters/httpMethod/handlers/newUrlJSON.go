package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	httpMethodHelpers "github.com/n0byk/short_url_backend/adapters/httpMethod/helpers"
	types "github.com/n0byk/short_url_backend/adapters/httpMethod/types"
	config "github.com/n0byk/short_url_backend/config"
	entities "github.com/n0byk/short_url_backend/dataservice/entities"
	filestorage "github.com/n0byk/short_url_backend/dataservice/filestorage"
	memory "github.com/n0byk/short_url_backend/dataservice/memory"
	helpers "github.com/n0byk/short_url_backend/helpers"
)

type props struct {
	URL string `json:"url"`
}

func NewURLJSON(w http.ResponseWriter, r *http.Request) {

	var urlBytes props

	err := json.NewDecoder(r.Body).Decode(&urlBytes)
	if err != nil {
		log.Print("ERR NewUrl - " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !httpMethodHelpers.ValidateURL(string(urlBytes.URL)) {
		log.Print("Validate error - " + string(urlBytes.URL))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token := helpers.GenerateToken(7)

	var adapter memory.URLCatalog

	if len(config.AppEnv().FileStoragePath) > 1 {
		storage, err := filestorage.NewStorageSet(config.AppEnv().FileStoragePath)
		if err != nil {
			log.Fatal(err)
		}
		defer storage.CloseURLCatalog()
		storage.WriteURL(&entities.URLCatalog{ShortURL: token, FullURL: string(urlBytes.URL)})
	} else {
		adapter.AddElement(entities.URLCatalog{ShortURL: token, FullURL: string(urlBytes.URL)})
	}

	person := types.JSONResponse{Result: config.AppEnv().BaseURL + "/" + token}
	response, jsonError := json.Marshal(person)

	if jsonError != nil {
		log.Print("Unable to encode JSON")
	}
	httpMethodHelpers.JSONResponse(w, response, http.StatusCreated)
}
