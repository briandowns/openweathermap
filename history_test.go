package openweathermap

import (
	"testing"
)

func validDataUnit(h *HistoricalWeatherData) bool {
	for _, m := range dataUnits {
		if h.Units == m {
			return true
		}
	}
	return false
}

func TestNewHistorical(t *testing.T) {
	h, err := NewHistorical("imperial")
	if err != nil {
		t.Error("Failed creating instance of HistoricalWeatherData")
	}
	if !validDataUnit(h) {
		t.Error("Invalid data unit")
	}
}
