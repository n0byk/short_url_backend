package handlers

import (
	"io/ioutil"
	"log"
	"net/http"

	Storage "github.com/n0byk/short_url_backend/httpapi/v1/storage"
	Entities "github.com/n0byk/short_url_backend/httpapi/v1/storage/entities"
	UrlCatalog "github.com/n0byk/short_url_backend/httpapi/v1/storage/repositories"
)

func NewUrl(w http.ResponseWriter, r *http.Request) {

	urlBytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Print("ERR NewUrl - " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	p := Entities.AddElement{
		ShortUrl: Storage.RandomToken(7),
		UrlBytes: urlBytes,
	}

	UrlCatalog.AddElement(p)
	w.Header().Set("Content-Type", "application/text")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("http://localhost:8080/" + p.ShortUrl))

}
