package version1

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gorilla/mux"
	"github.com/n0byk/short_url_backend/config"
	"github.com/n0byk/short_url_backend/dataservice/memory"
)

// This example demonstrates how to add new url to the system
func ExampleAddURLHandler() {
	r := mux.NewRouter()
	storage := memory.NewMemoryRepository()
	config.AppService = config.Service{ShortLinkLen: 7, BaseURL: "http://localhost:8080", Storage: storage}
	route := r.NewRoute().HeadersRegexp("Accept", "text/plain")

	req1, _ := http.NewRequest("GET", "http://ya.ru", nil)
	req1.Header.Add("Accept", "text/plain")

	req2, _ := http.NewRequest("GET", "http://ya123.ru", nil)
	req2.Header.Set("Accept", "text/plain")

	matchInfo := &mux.RouteMatch{}
	fmt.Printf("Match: %v %q\n", route.Match(req1, matchInfo), req1.Header["Accept"])
	fmt.Printf("Match: %v %q\n", route.Match(req2, matchInfo), req2.Header["Accept"])
	// Output:
	// Match: true ["text/plain"]
	// Match: true ["text/plain"]
}

// This example demonstrates how to add new url to the system
func ExampleDBPing() {

	r := mux.NewRouter()
	storage := memory.NewMemoryRepository()
	config.AppService = config.Service{ShortLinkLen: 7, BaseURL: "http://localhost:8080", Storage: storage}
	route := r.NewRoute().HeadersRegexp("Accept", "text/plain")

	req := httptest.NewRequest("GET", "http://localhost:8080/ping", nil)
	matchInfo := &mux.RouteMatch{}
	fmt.Printf("Match: %v %q\n", route.Match(req, matchInfo), req.Header["Accept"])

	// Output:
	// Match: false []
}

func TestRouters(t *testing.T) {
	tests := []struct {
		name string
		want *mux.Router
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Routers(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Routers() = %v, want %v", got, tt.want)
			}
		})
	}
}
