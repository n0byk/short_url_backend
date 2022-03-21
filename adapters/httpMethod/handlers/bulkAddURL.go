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
	Correlation_id string `json:"correlation_id"`
	Original_url   string `json:"original_url"`
}

func BulkAddURL(w http.ResponseWriter, r *http.Request) {

	var urlBytes []bulkProps
	bodyBytes, err := httpMethodHelpers.ReadBodyBytes(r)
	if err != nil {
		log.Println("Error while getting body - ")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// userID, err := r.Cookie("user_id")
	// if err != nil {
	// 	log.Println("Can't get user_id ")
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	err = json.Unmarshal(bodyBytes, &urlBytes)
	if err != nil {
		log.Print("ERR NewUrl - " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// if !httpMethodHelpers.ValidateURL(string(urlBytes.URL)) {
	// 	log.Print("Validate error - " + string(urlBytes.URL))
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	for i, bp := range urlBytes {
		fmt.Println(i)
		fmt.Println(bp)

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
