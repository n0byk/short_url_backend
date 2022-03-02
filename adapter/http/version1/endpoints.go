package version1

import (
	"github.com/gorilla/mux"

	handlers "github.com/n0byk/short_url_backend/adapter/http/version1/handlers"
)

func V1() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/", handlers.NewURL).Methods("POST")
	router.HandleFunc("/{url}", handlers.GetURL).Methods("GET")
	router.HandleFunc("/api/shorten", handlers.NewURLJSON).Methods("POST")

	return router
}
