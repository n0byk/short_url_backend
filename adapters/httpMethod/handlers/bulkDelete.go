package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	config "github.com/n0byk/short_url_backend/config"
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

	// go func() {
	// 	log.Println("done")
	// 	config.AppService.Storage.BulkDelete(ids, userID.Value)
	// }()

	errC := make(chan error)
	go func() {
		log.Println("done")
		errC <- config.AppService.Storage.BulkDelete(ids, userID.Value)
	}()

	err = <-errC

	if err != nil {
		log.Println("Can't BulkDelete ", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(nil)
		return
	}

	log.Println("responce")
	w.WriteHeader(http.StatusAccepted)
}
