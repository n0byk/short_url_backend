package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	config "github.com/n0byk/short_url_backend/config"
	"golang.org/x/sync/errgroup"
)

func BulkDelete(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		log.Println("Can't get body ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var ids []string
	err = json.Unmarshal(body, &ids)
	if err != nil {
		log.Println("Can't get json body ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(ids)

	userID, err := r.Cookie("user_id")
	if err != nil {
		log.Println("Can't get user_id ")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(nil)
		return
	}
	g, _ := errgroup.WithContext(r.Context())

	g.Go(func() error {
		return config.AppService.Storage.BulkDelete(ids, userID.Value)
	})

	if err = g.Wait(); err != nil {
		log.Println("Can't delete items ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Println("responce")
	w.WriteHeader(http.StatusAccepted)
}
