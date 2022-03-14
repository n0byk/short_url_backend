package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	httpMethodHelpers "github.com/n0byk/short_url_backend/adapters/httpMethod/helpers"
	types "github.com/n0byk/short_url_backend/adapters/httpMethod/types"
	config "github.com/n0byk/short_url_backend/config"
	helpers "github.com/n0byk/short_url_backend/helpers"
)

type props struct {
	URL string `json:"url"`
}

func AddURLJSON(w http.ResponseWriter, r *http.Request) {

	var urlBytes props
	bodyBytes, err := httpMethodHelpers.ReadBodyBytes(r)
	if err != nil {
		log.Println("Error while getting body - ")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(bodyBytes, &urlBytes)
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

	token := helpers.GenerateToken(config.AppService.ShortLinkLen)

	config.AppService.Storage.AddURL(token, string(urlBytes.URL))

	person := types.JSONResponse{Result: config.AppService.BaseURL + "/" + token}
	response, jsonError := json.Marshal(person)

	if jsonError != nil {
		log.Print("Unable to encode JSON")
	}
	httpMethodHelpers.JSONResponse(w, response, http.StatusCreated)
}
