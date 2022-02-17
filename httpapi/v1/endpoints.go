package v1

import (
	"github.com/gorilla/mux"

	handlers "github.com/n0byk/short_url_backend/httpapi/v1/handlers"
)

func V1() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/", handlers.NewUrl).Methods("POST")
	router.HandleFunc("/{id}", handlers.GetUrl).Methods("GET")

	return router
}
