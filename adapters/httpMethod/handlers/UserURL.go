package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	httpMethodHelpers "github.com/n0byk/short_url_backend/adapters/httpMethod/helpers"
	config "github.com/n0byk/short_url_backend/config"
)

// Get user added urls handler
func UserURL(w http.ResponseWriter, r *http.Request) {
	userID, err := r.Cookie("user_id")
	if err != nil {
		log.Println("Can't get user_id ")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data, _ := config.AppService.Storage.GetUserData(context.Background(), userID.Value)

	response, jsonError := json.Marshal(data)

	if jsonError != nil {
		log.Print("Unable to encode JSON")
	}

	if len(data) > 0 {
		httpMethodHelpers.JSONResponse(w, response, http.StatusOK)
		return
	}

	httpMethodHelpers.JSONResponse(w, response, http.StatusNoContent)
}
