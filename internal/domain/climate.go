package domain

import (
	"clima-cep/internal/outbound"
	"errors"
)

var ErrZipcodeNotFound = errors.New("zipcode not found")

type Climate struct {
	TempC float64
	TempF float64
	TempK float64
}

type ClimateService struct {
	outboundService outbound.Service
}

func NewClimateService(service outbound.Service) *ClimateService {
	return &ClimateService{outboundService: service}
}

func (s *ClimateService) GetClimate(zipcode string) (*Climate, error) {
	location, err := outbound.FetchLocationByZipcode(zipcode)
	if err != nil {
		return nil, ErrZipcodeNotFound
	}

	tempC, err := outbound.FetchWeather(location)
	if err != nil {
		return nil, errors.New("failed to fetch weather data")
	}

	return &Climate{
		TempC: tempC,
		TempF: tempC*1.8 + 32,
		TempK: tempC + 273,
	}, nil
}
