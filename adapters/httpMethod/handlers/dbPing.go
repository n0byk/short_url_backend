package handlers

import (
	"net/http"
)

func DBPing(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("200"))
}
