package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	config "github.com/n0byk/short_url_backend/config"
)

var wg sync.WaitGroup

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

	wg.Add(1)
	go func() {
		defer wg.Done()
		config.AppService.Storage.BulkDelete(r.Context(), ids, userID.Value)
	}()

	w.WriteHeader(http.StatusAccepted)
}
