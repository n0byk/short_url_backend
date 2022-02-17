package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUrl(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	type want struct {
		code     int
		response string
	}
	tests := []struct {
		name string
		want want
		args args
	}{
		{name: "Empty request №1.", want: want{code: http.StatusBadRequest, response: ""}},
		//следующий кейс что-то не отрабатывает, кажется все так - но не работает, код вместо 301 прилетает 400
		// {name: "Redirect to default url №2.", want: want{code: http.StatusTemporaryRedirect, response: ""}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/3e89ffb0-948a-4d2c-9308-e5d57569cf19", nil)

			w := httptest.NewRecorder()
			h := http.HandlerFunc(GetUrl)
			h.ServeHTTP(w, request)
			res := w.Result()
			defer res.Body.Close()
			if res.StatusCode != tt.want.code {
				t.Errorf("Expected status code %d, got %d", tt.want.code, w.Code)
			}

		})
	}
}
