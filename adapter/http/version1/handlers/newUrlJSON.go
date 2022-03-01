package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	types "github.com/n0byk/short_url_backend/adapter/http/version1/types"
	entities "github.com/n0byk/short_url_backend/dataservice/entities"
	mockdata "github.com/n0byk/short_url_backend/dataservice/mockdata"
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

	if !helpers.ValidateURL(string(urlBytes.URL)) {
		log.Print("Validate error - " + string(urlBytes.URL))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token := helpers.GenerateToken(7)

	var adapter mockdata.URLCatalog

	adapter.AddElement(entities.URLCatalog{ShortURL: token, FullURL: string(urlBytes.URL)})

	person := types.JSONResponse{Result: "http://localhost:8080/" + token}

	response, jsonError := json.Marshal(person)

	if jsonError != nil {
		log.Print("Unable to encode JSON")
	}

	helpers.JSONResponse(w, response, http.StatusCreated)
}
