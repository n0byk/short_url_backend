package handlers

import (
	"log"
	"net/http"

	httpMethodHelpers "github.com/n0byk/short_url_backend/adapters/httpMethod/helpers"
	config "github.com/n0byk/short_url_backend/config"
)

func AddURLHandler(w http.ResponseWriter, r *http.Request) {

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

	if !httpMethodHelpers.ValidateURL(string(bodyBytes)) {
		log.Println("Validate error - " + string(bodyBytes))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := config.AppService.Storage.AddURL(string(bodyBytes), userID.Value)
	if err != nil {
		log.Println("Some error while adding new URL")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var data = []byte(config.AppService.BaseURL + "/" + token)
	config.AppService.Storage.SetUserData(string(bodyBytes), string(data), userID.Value)

	w.Header().Set("Content-Type", "application/text; charset=utf-8")

	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}
