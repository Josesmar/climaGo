package outbound

import "github.com/stretchr/testify/mock"

// Service é a interface que o ClimateService usa para interagir com o pacote outbound
type Service interface {
	FetchLocationByZipcode(zipcode string) (string, error)
	FetchWeather(location string) (float64, error)
}

// MockOutboundService é o mock da interface Service
type MockOutboundService struct {
	mock.Mock
}

// FetchLocationByZipcode é o método mockado para FetchLocationByZipcode
func (m *MockOutboundService) FetchLocationByZipcode(zipcode string) (string, error) {
	args := m.Called(zipcode)
	return args.String(0), args.Error(1)
}

// FetchWeather é o método mockado para FetchWeather
func (m *MockOutboundService) FetchWeather(location string) (float64, error) {
	args := m.Called(location)
	return args.Get(0).(float64), args.Error(1)
}
