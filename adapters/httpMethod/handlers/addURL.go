package handlers

import (
	"fmt"
	"log"
	"net/http"

	httpMethodHelpers "github.com/n0byk/short_url_backend/adapters/httpMethod/helpers"
	config "github.com/n0byk/short_url_backend/config"
)

func AddURLHandler(w http.ResponseWriter, r *http.Request) {
fmt.Println("AddURLHandler")
	bodyBytes, err := httpMethodHelpers.ReadBodyBytes(r)

	if err != nil {
		log.Println("Error while getting body - ")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
		return
	}

	userID, err := r.Cookie("user_id")
	if err != nil {

		log.Println("Can't get user_id ")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(nil)
		return
	}

	if !httpMethodHelpers.ValidateURL(string(bodyBytes)) {
		log.Println("Validate error - " + string(bodyBytes))

		w.WriteHeader(http.StatusBadRequest)
		w.Write(nil)
		return
	}

	token, duplicate, err := config.AppService.Storage.AddURL(string(bodyBytes), userID.Value)
	if err != nil {
		log.Println("Some error while adding new URL")
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var data = []byte(config.AppService.BaseURL + "/" + token)
	config.AppService.Storage.SetUserData(string(bodyBytes), string(data), userID.Value)
	fmt.Println("addURL" + string(bodyBytes))
	w.Header().Set("Content-Type", "application/text; charset=utf-8")
	if duplicate {
		w.WriteHeader(http.StatusConflict)
	} else {
		w.WriteHeader(http.StatusCreated)
	}

	w.Write(data)
}
