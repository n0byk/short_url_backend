package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	httpMethodHelpers "github.com/n0byk/short_url_backend/adapters/httpMethod/helpers"
	config "github.com/n0byk/short_url_backend/config"
)

func UserURL(w http.ResponseWriter, r *http.Request) {
	// httpMethodHelpers.SetCookie(w, r)
	userID, err := r.Cookie("user_id")
	if err != nil {
		log.Println("Can't get user_id ")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := config.AppService.Storage.GetUserData(userID.Value)
	if err != nil {
		log.Println("Err ", err)
		w.WriteHeader(http.StatusNoContent)
	}
	response, jsonError := json.Marshal(data)

	if jsonError != nil {
		log.Print("Unable to encode JSON")
	}

	if len(data) > 0 {
		w.WriteHeader(http.StatusOK)
		httpMethodHelpers.JSONResponse(w, response)
		return
	}
	w.WriteHeader(http.StatusNoContent)

	httpMethodHelpers.JSONResponse(w, response)
}
