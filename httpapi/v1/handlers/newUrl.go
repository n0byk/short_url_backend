package handlers

import (
	"io/ioutil"
	"log"
	"net/http"

	helpers "github.com/n0byk/short_url_backend/httpapi/v1/helpers"
	entities "github.com/n0byk/short_url_backend/httpapi/v1/storage/entities"
	urlCatalog "github.com/n0byk/short_url_backend/httpapi/v1/storage/repositories"
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

	p := entities.AddElement{
		ShortURL: helpers.GenerateToken(7),
		URLBytes: urlBytes,
	}

	urlCatalog.AddElement(p)
	w.Header().Set("Content-Type", "application/text")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("http://localhost:8080/" + p.ShortURL))

}
