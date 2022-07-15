package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"net/http"

	httpMethodHelpers "github.com/n0byk/short_url_backend/adapters/httpMethod/helpers"
	config "github.com/n0byk/short_url_backend/config"
)

// Get stat info handler
func Stats(w http.ResponseWriter, r *http.Request) {
	realIP := r.Header.Get("X-Real-IP")
	if config.AppService.Env.TrustedCIDR != nil && config.AppService.Env.TrustedCIDR.Contains(net.ParseIP(realIP)) {
		log.Println("X-Real-IP")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	data, _ := config.AppService.Storage.StatInfo(context.Background())

	response, jsonError := json.Marshal(data)

	if jsonError != nil {
		log.Print("Unable to encode JSON")
	}

	httpMethodHelpers.JSONResponse(w, response, http.StatusOK)
}
