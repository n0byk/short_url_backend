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

type bulkProps struct {
	CorrelationID string `json:"correlation_id"`
	OriginalURL   string `json:"original_url"`
}

func BulkAddURL(w http.ResponseWriter, r *http.Request) {
	fmt.Println("BulkAddURL")
	var urlBytes []bulkProps
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
	var person []types.BulkJSONResponse

	for _, item := range urlBytes {
		if !httpMethodHelpers.ValidateURL(item.OriginalURL) {
			log.Print("Validate error - " + item.OriginalURL)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		token, _, err := config.AppService.Storage.AddURL(string(item.OriginalURL), userID.Value)
		if err != nil {
			log.Println("Unable to get token")
		}

		config.AppService.Storage.SetUserData(string(bodyBytes), config.AppService.BaseURL+"/"+token, userID.Value)
		person = append(person, types.BulkJSONResponse{ShortURL: config.AppService.BaseURL + "/" + token, CorrelationID: item.CorrelationID})
	}

	response, jsonError := json.Marshal(person)
	if jsonError != nil {
		log.Println("Unable to encode JSON")
	}

	httpMethodHelpers.JSONResponse(w, response, http.StatusCreated)
}
