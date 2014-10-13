package openweathermap

import (
	"errors"
	"strings"
)

// NewHistorical returns a new HistoricalWeatherData pointer with the supplied
// arguments.
func NewHistorical(unit string) (*HistoricalWeatherData, error) {
	unitChoice := strings.ToLower(unit)
	for _, i := range dataUnits {
		if strings.Contains(unitChoice, i) {
			return &HistoricalWeatherData{Units: unitChoice}, nil
		}
	}
	return nil, errors.New("ERROR: unit of measure not available")
}
