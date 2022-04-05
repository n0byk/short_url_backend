package version1

import (
	"github.com/gorilla/mux"

	handlers "github.com/n0byk/short_url_backend/adapters/httpMethod/handlers"
)

func Routers() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/shorten/batch", handlers.BulkAddURL).Methods("POST")
	router.HandleFunc("/api/user/urls", handlers.BulkDelete).Methods("DELETE")
	router.HandleFunc("/ping", handlers.DBPing).Methods("GET")
	router.HandleFunc("/", handlers.AddURLHandler).Methods("POST")
	router.HandleFunc("/{url}", handlers.GetURL).Methods("GET")
	router.HandleFunc("/api/shorten", handlers.AddURLJSON).Methods("POST")
	router.HandleFunc("/api/user/urls", handlers.UserURL).Methods("GET")

	return router
}
