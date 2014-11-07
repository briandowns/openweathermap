package openweathermap

import (
	"testing"
)

func TestNewForecast(t *testing.T) {
	f, err := NewForecast("imperial")
	if err != nil {
		t.Error("Failed to setup ForecastWeatherData")
	}
	if !ValidDataUnit(f.Units) {
		t.Error("Failed creating instance of HistoricalWeatherData")
	}
}
