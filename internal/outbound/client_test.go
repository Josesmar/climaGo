package outbound

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
)

func Test_FetchLocationByZipcode(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/ws/12345678/json/", nil)
	resp := httptest.NewRecorder()

	// Simulando resposta de ViaCEP
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"localidade": "São Paulo"}`))
	})
	handler.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "São Paulo")
}

func Test_FetchWeather(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/weather/current.json?key=apiKey&q=Sao%20Paulo", nil)
	resp := httptest.NewRecorder()

	// Simulando resposta de WeatherAPI
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"current": {"temp_c": 25.5}}`))
	})
	handler.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), `"temp_c": 25.5`)
}
