package openweathermap

import (
	"testing"
)

// TestPollutionByParams tests the call to the pollution API
func TestPollutionByParams(t *testing.T) {
	t.Parallel()
	p := NewPollution()
	params := &PollutionParameters{
		Location: Coordinates{
			Latitude:  0.0,
			Longitude: 10.0,
		},
		Datetime: "current",
	}
	if err := p.PollutionByParams(params); err != nil {
		t.Error(err)
	}
}
