package openweathermap

import (
	"errors"
	"strings"
)

type ForecastWeatherData struct {
}

// NewHistorical returns a new HistoricalWeatherData pointer with the supplied
// arguments.
func NewForecast(unit string) (*ForecastWeatherData, error) {
	unitChoice := strings.ToLower(unit)
	for _, i := range dataUnits {
		if strings.Contains(unitChoice, i) {
			return &ForecastWeatherData{Units: unitChoice}, nil
		}
	}
	return nil, errors.New("ERROR: unit of measure not available")
}
