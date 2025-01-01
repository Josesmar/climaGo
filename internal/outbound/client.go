package outbound

import (
	"clima-cep/internal/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type WeatherAPIResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
	Localidade string `json:"localidade"`
}

func FetchLocationByZipcode(zipcode string) (string, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", zipcode)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error calling ViaCEP: %v", err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("ViaCEP returned status: %d", resp.StatusCode)
		return "", fmt.Errorf("failed to fetch location")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading ViaCEP response: %v", err)
		return "", err
	}

	log.Printf("ViaCEP response: %s", body)
	var weatherAPIResponse WeatherAPIResponse
	err = json.Unmarshal(body, &weatherAPIResponse)
	if err != nil {
		log.Printf("Error unmarshalling JSON: %v", err)
		return "", err
	}
	return weatherAPIResponse.Localidade, nil

}

func FetchWeather(location string) (float64, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Printf("Error loading configuration: %v", err)
		return 0, err
	}

	if cfg.WeatherAPIKey == "" {
		log.Printf("Error: WEATHERAPI_KEY is missing")
		return 0, fmt.Errorf("missing WeatherAPI key")
	}

	encodedLocation := url.QueryEscape(location)

	url := fmt.Sprintf("%s/current.json?key=%s&q=%s", cfg.WeatherAPIBaseURL, cfg.WeatherAPIKey, encodedLocation)
	log.Printf("Fetching weather data from: %s", url)

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error calling WeatherAPI: %v", err)
		return 0, err
	}
	defer resp.Body.Close()

	log.Printf("WeatherAPI status code: %d", resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		log.Printf("WeatherAPI returned status: %d", resp.StatusCode)
		return 0, fmt.Errorf("failed to fetch weather data, status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading WeatherAPI response: %v", err)
		return 0, err
	}

	log.Printf("WeatherAPI response body: %s", body)

	var weatherAPIResponse WeatherAPIResponse
	err = json.Unmarshal(body, &weatherAPIResponse)
	if err != nil {
		log.Printf("Error unmarshalling JSON: %v", err)
		return 0, err
	}
	return weatherAPIResponse.Current.TempC, nil

}
