package mocks

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

func NewPaymentMakerServer() *httptest.Server {

	response := CreatePaymentStartResponseMock()
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/holidays":
			if r.Method != http.MethodGet {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(response)

		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))

	println("Mock server running at:", mockServer.URL)
	return mockServer
}
