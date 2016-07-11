package openweathermap

import (
	"net/http"
	"reflect"
	"testing"
	"time"
)

// TestNewPollution
func TestNewPollution(t *testing.T) {

	p, err := NewPollution()
	if err != nil {
		t.Error(err)
	}

	if reflect.TypeOf(p).String() != "*openweathermap.Pollution" {
		t.Error("incorrect data type returned")
	}
}

// TestNewPollution with custom http client
func TestNewPollutionWithCustomHttpClient(t *testing.T) {

	hc := http.DefaultClient
	hc.Timeout = time.Duration(1) * time.Second
	p, err := NewPollution(WithHttpClient(hc))
	if err != nil {
		t.Error(err)
	}

	if reflect.TypeOf(p).String() != "*openweathermap.Pollution" {
		t.Error("incorrect data type returned")
	}

	expected := time.Duration(1) * time.Second
	if p.client.Timeout != expected {
		t.Errorf("Expected Duration %v, but got %v", expected, p.client.Timeout)
	}
}

func TestValidAlias(t *testing.T) {
	t.Parallel()
	testAliases := []string{"now", "then", "current"}
	for _, i := range testAliases {
		if !ValidAlias(i) {
			t.Log("received expected failure")
		}
	}
}

// TestPollutionByParams tests the call to the pollution API
func TestPollutionByParams(t *testing.T) {
	t.Parallel()
	p, err := NewPollution()
	if err != nil {
		t.Error(err)
	}
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
