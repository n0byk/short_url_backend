package v1

import (
	"github.com/gorilla/mux"

	handler "github.com/n0byk/short_url_backend/httpapi/v1/handlers"
)

func V1() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/", handler.NewUrl).Methods("POST")
	router.HandleFunc("/{id}", handler.GetUrl).Methods("GET")

	return router
}
