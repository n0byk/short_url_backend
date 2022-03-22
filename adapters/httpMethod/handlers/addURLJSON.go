package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	httpMethodHelpers "github.com/n0byk/short_url_backend/adapters/httpMethod/helpers"
	types "github.com/n0byk/short_url_backend/adapters/httpMethod/types"
	config "github.com/n0byk/short_url_backend/config"
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

	userID, err := r.Cookie("user_id")
	if err != nil {
		log.Println("Can't get user_id ")
		http.Error(w, err.Error(), http.StatusBadRequest)
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

	token, duplicate, err := config.AppService.Storage.AddURL(string(urlBytes.URL), userID.Value)
	if err != nil {
		log.Print("Unable to get token")
	}
	config.AppService.Storage.SetUserData(string(bodyBytes), config.AppService.BaseURL+"/"+token, userID.Value)

	person := types.JSONResponse{Result: config.AppService.BaseURL + "/" + token}
	response, jsonError := json.Marshal(person)

	if jsonError != nil {
		log.Print("Unable to encode JSON")
	}

	fmt.Println("addURLJSON" + string(bodyBytes))
	if duplicate {
		httpMethodHelpers.JSONResponse(w, response, http.StatusConflict)
		return
	}
	httpMethodHelpers.JSONResponse(w, response, http.StatusCreated)
}
