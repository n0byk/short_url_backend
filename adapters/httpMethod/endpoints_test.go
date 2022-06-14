package version1

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/gorilla/mux"
	"github.com/n0byk/short_url_backend/config"
	"github.com/n0byk/short_url_backend/dataservice/memory"
)

// This example demonstrates how to add new url to the system
func ExampleAddURLHandler_Post() {
	AddURLHandler{}.Post() // Add new url to the system
	// Output:
	// # Request
	// POST http://localhost:8081/  HTTP/1.1
	// http://ya.ry/
	// Capsule-Version: 1.2.3
	//
	// # Response
	// HTTP/1.1 200 OK
	//
	// http://localhost:8081/VzHAYQi
}

// This endpoint check DB connection
func ExampleDBPing_Get() {
	DBPing{}.Get() // Check DB connection
	// Output:
	// # Request
	// GET http://localhost:8080/ping HTTP/1.1
	// Capsule-Version: 1.2.3
	//
	// # Response
	// HTTP/1.1 200 OK
	// Access-Control-Allow-Origin: https://www.capsulerx.com
	// Content-Type: application/vnd.api+json
}
