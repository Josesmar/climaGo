package config

import (
	"bufio"
	"log"
	"os"
	"strings"

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

	weatherAPIKey, err := readAPIKeyFromFile("secret_key.txt")
	if err != nil {
		log.Println("Error reading API key: ", err)
		return nil, err
	}

	config := &Config{
		WeatherAPIBaseURL: os.Getenv("WEATHER_API_BASE_URL"),
		ViaCepBaseURL:     os.Getenv("VIA_CEP_BASE_URL"),
		WeatherAPIKey:     weatherAPIKey,
	}

	return config, err
}

func readAPIKeyFromFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "WEATHER_API_KEY=") {
			return strings.TrimPrefix(line, "WEATHER_API_KEY="), nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", nil
}
