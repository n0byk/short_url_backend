package handlers

import (
	"log"
	"net/http"

	config "github.com/n0byk/short_url_backend/config"
)

func BulkDelete(w http.ResponseWriter, r *http.Request) {
	userID, err := r.Cookie("user_id")
	if err != nil {

		log.Println("Can't get user_id ")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(nil)
		return
	}
	log.Println(r)

	go func() {
		log.Println("done")
		config.AppService.Storage.BulkDelete(r.URL.Query(), userID.Value)
	}()

	log.Println("responce")
	w.WriteHeader(http.StatusAccepted)
}
