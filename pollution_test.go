package openweathermap

import (
	"net/http"
	"os"
	"reflect"
	"testing"
	"time"
)

// TestNewPollution
func TestNewPollution(t *testing.T) {

	p, err := NewPollution(os.Getenv("OWM_API_KEY"))
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
	p, err := NewPollution(os.Getenv("OWM_API_KEY"), WithHttpClient(hc))
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

// TestNewPollutionWithInvalidOptions will verify that returns an error with
// invalid option
func TestNewPollutionWithInvalidOptions(t *testing.T) {

	optionsPattern := [][]Option{
		{nil},
		{nil, nil},
		{WithHttpClient(&http.Client{}), nil},
		{nil, WithHttpClient(&http.Client{})},
	}

	for _, options := range optionsPattern {
		c, err := NewPollution(os.Getenv("OWM_API_KEY"), options...)
		if err == errInvalidOption {
			t.Logf("Received expected invalid option error. message: %s", err.Error())
		} else if err != nil {
			t.Errorf("Expected %v, but got %v", errInvalidOption, err)
		}
		if c != nil {
			t.Errorf("Expected nil, but got %v", c)
		}
	}
}

// TestNewPollutionWithInvalidHttpClient will verify that returns an error with
// invalid http client
func TestNewPollutionWithInvalidHttpClient(t *testing.T) {

	p, err := NewPollution(os.Getenv("OWM_API_KEY"), WithHttpClient(nil))
	if err == errInvalidHttpClient {
		t.Logf("Received expected bad client error. message: %s", err.Error())
	} else if err != nil {
		t.Errorf("Expected %v, but got %v", errInvalidHttpClient, err)
	}
	if p != nil {
		t.Errorf("Expected nil, but got %v", p)
	}
}

// TestNewPollutionWithApiURL  will verify that a new instance of PollutionWeatherData
// is created with custom API url
func TestNewPollutionWithApiURL(t *testing.T) {
	c, err := NewPollution(os.Getenv("OWM_API_KEY"), WithApiURL("https://ru.openweathermap.org/"))
	if err != nil {
		t.Error(err)
	}

	if reflect.TypeOf(c).String() != "*openweathermap.Pollution" {
		t.Error("incorrect data type returned")
	}

	expected := "https://ru.openweathermap.org/pollution/v1/co/"
	got := c.pollutionURL
	if got != expected {
		t.Errorf("Expected pollutionURL %v, but got %v", expected, got)
	}
}

// TestNewPollutionWithInvalidApiURL  will verify that returns an error with
// invalid API url
func TestNewPollutionWithInvalidApiURL(t *testing.T) {
	c, err := NewPollution(os.Getenv("OWM_API_KEY"), WithApiURL("somestring"))
	if err == errInvalidApiURL {
		t.Logf("Received expected invalid url error. message: %s", err.Error())
	} else if err != nil {
		t.Errorf("Expected %v, but got %v", errInvalidApiURL, err)
	}
	if c != nil {
		t.Errorf("Expected nil, but got %v", c)
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
	p, err := NewPollution(os.Getenv("OWM_API_KEY"))
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
