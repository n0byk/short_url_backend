package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	httpMethodHelpers "github.com/n0byk/short_url_backend/adapters/httpMethod/helpers"
	config "github.com/n0byk/short_url_backend/config"
	helpers "github.com/n0byk/short_url_backend/helpers"
)

type bulkProps struct {
	CorrelationID string `json:"correlation_id"`
	OriginalURL   string `json:"original_url"`
}

func BulkAddURL(w http.ResponseWriter, r *http.Request) {

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
	fmt.Println(userID)
	err = json.Unmarshal(bodyBytes, &urlBytes)
	if err != nil {
		log.Print("ERR NewUrl - " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, item := range urlBytes { 
		if !httpMethodHelpers.ValidateURL(item.OriginalURL) {
			log.Print("Validate error - " + item.OriginalURL)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

	}

	token := helpers.GenerateToken(config.AppService.ShortLinkLen)

	// config.AppService.Storage.AddURL(token, string(urlBytes.URL), userID.Value)

	// config.AppService.Storage.SetUserData(string(bodyBytes), config.AppService.BaseURL+"/"+token, userID.Value)

	// person := types.JSONResponse{Result: config.AppService.BaseURL + "/" + token}
	// response, jsonError := json.Marshal(person)

	// if jsonError != nil {
	// 	log.Print("Unable to encode JSON")
	// }
	httpMethodHelpers.JSONResponse(w, []byte(token), http.StatusCreated)
}
