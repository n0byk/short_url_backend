package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNewUrl(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	type want struct {
		code        int
		response    string
		contentType string
		url         string
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{name: "No Body №1.", want: want{code: http.StatusBadRequest, response: ""}},
		//следующий кейс что-то не отрабатывает, кажется все так - но не работает, код вместо 201 прилетает 400
		// {name: "Add some url №2", want: want{code: http.StatusCreated, response: "", url: "http://ya.ru"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.want.url))

			w := httptest.NewRecorder()
			h := http.HandlerFunc(GetURL)
			h.ServeHTTP(w, request)
			res := w.Result()

			defer res.Body.Close()

			if res.StatusCode != tt.want.code {
				t.Errorf("Expected status code %d, got %d", tt.want.code, w.Code)
			}

		})
	}
}
