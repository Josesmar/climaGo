package domain

import (
	"errors"
	"testing"

	"clima-cep/internal/outbound"

	"github.com/stretchr/testify/assert"
)

func Test_GetClimate(t *testing.T) {
	mockOutbound := new(outbound.MockOutboundService)

	mockOutbound.On("FetchLocationByZipcode", "12345678").Return("Localidade", nil)
	mockOutbound.On("FetchWeather", "Localidade").Return(25.5, nil)

	service := NewClimateService(mockOutbound)

	climate, err := service.GetClimate("12345678")
	assert.NoError(t, err)
	assert.NotNil(t, climate)
	assert.Equal(t, 25.5, climate.TempC)
	assert.Equal(t, 77.9, climate.TempF)
	assert.Equal(t, 298.65, climate.TempK)

	mockOutbound.AssertExpectations(t)

	mockOutbound.On("FetchLocationByZipcode", "87654321").Return("", errors.New("zipcode not found"))

	climate, err = service.GetClimate("87654321")
	assert.Error(t, err)
	assert.Nil(t, climate)

	mockOutbound.AssertExpectations(t)
}
