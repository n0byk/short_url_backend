package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	httpMethodHelpers "github.com/n0byk/short_url_backend/adapters/httpMethod/helpers"
	config "github.com/n0byk/short_url_backend/config"
	helpers "github.com/n0byk/short_url_backend/helpers"
)

func AddURLHandler(w http.ResponseWriter, r *http.Request) {
	urlBytes, err := ioutil.ReadAll(r.Body)
	fmt.Print(string(urlBytes))
	if err != nil {
		log.Println("ERR NewUrl - " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !httpMethodHelpers.ValidateURL(string(urlBytes)) {
		log.Println("Validate error - " + string(urlBytes))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token := helpers.GenerateToken(config.AppService.ShortLinkLen)

	err = config.AppService.Storage.AddURL(token, string(urlBytes))
	if err != nil {
		log.Println("Some error while adding new URL")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data, _ := httpMethodHelpers.Compress([]byte(config.AppService.BaseURL + "/" + token))
	fmt.Println(data)

	// w.Header().Set("Content-Type", "application/text")
	w.Header().Set("Content-Encoding", "gzip")
	r.Header.Add("Accept-Encoding", "gzip")
	w.WriteHeader(http.StatusCreated)

	w.Write(data)
}
