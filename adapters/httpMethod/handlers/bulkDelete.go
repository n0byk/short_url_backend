package handlers

import (
	"fmt"
	"net/http"
)

func BulkDelete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello! Parameters: %v", r.URL.Query())

	// httpMethodHelpers.JSONResponse(w, response, http.StatusCreated)
}
