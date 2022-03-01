package handlers

import (
	"io/ioutil"
	"log"
	"net/http"

	entities "github.com/n0byk/short_url_backend/dataservice/entities"
	mockdata "github.com/n0byk/short_url_backend/dataservice/mockdata"
	helpers "github.com/n0byk/short_url_backend/helpers"
)

func NewURL(w http.ResponseWriter, r *http.Request) {

	urlBytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Print("ERR NewUrl - " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !helpers.ValidateURL(string(urlBytes)) {
		log.Print("Validate error - " + string(urlBytes))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token := helpers.GenerateToken(7)

	var adapter mockdata.URLCatalog

	adapter.AddElement(entities.URLCatalog{ShortURL: token, FullURL: string(urlBytes)})

	w.Header().Set("Content-Type", "application/text")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("http://localhost:8080/" + token))

}
