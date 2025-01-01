package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	WeatherAPIBaseURL string
	ViaCepBaseURL     string
	WeatherAPIKey     string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading config file: ", err)
	}

	config := &Config{
		WeatherAPIBaseURL: os.Getenv("WEATHER_API_BASE_URL"),
		ViaCepBaseURL:     os.Getenv("VIA_CEP_BASE_URL"),
		WeatherAPIKey:     os.Getenv("WEATHER_API_KEY"),
	}

	return config, err
}
