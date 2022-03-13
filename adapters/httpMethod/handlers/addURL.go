package handlers

import (
	"log"
	"net/http"

	httpMethodHelpers "github.com/n0byk/short_url_backend/adapters/httpMethod/helpers"
	config "github.com/n0byk/short_url_backend/config"
	helpers "github.com/n0byk/short_url_backend/helpers"
)

func AddURLHandler(w http.ResponseWriter, r *http.Request) {

	url, err := httpMethodHelpers.ReadBodyBytes(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !httpMethodHelpers.ValidateURL(string(url)) {
		log.Println("Validate error - " + string(url))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token := helpers.GenerateToken(config.AppService.ShortLinkLen)

	err = config.AppService.Storage.AddURL(token, string(url))
	if err != nil {
		log.Println("Some error while adding new URL")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var data []byte = []byte(config.AppService.BaseURL + "/" + token)

	if r.Header.Get("Accept-Encoding") == "gzip" {
		data, _ = httpMethodHelpers.Compress(data)
		w.Header().Set("Content-Encoding", "gzip")
	} else {
		w.Header().Set("Content-Type", "application/text")
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}
