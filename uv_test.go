package openweathermap

import (
	"net/http"
	"os"
	"reflect"
	"testing"
	"time"
)

var coords = &Coordinates{
	Longitude: 53.343497,
	Latitude:  -6.288379,
}

// TestNewUV
func TestNewUV(t *testing.T) {

	uv, err := NewUV(os.Getenv("OWM_API_KEY"))
	if err != nil {
		t.Error(err)
	}

	if reflect.TypeOf(uv).String() != "*openweathermap.UV" {
		t.Error("incorrect data type returned")
	}
}

// TestNewUV with custom http client
func TestNewUVWithCustomHttpClient(t *testing.T) {

	hc := http.DefaultClient
	hc.Timeout = time.Duration(1) * time.Second
	uv, err := NewUV(os.Getenv("OWM_API_KEY"), WithHttpClient(hc))
	if err != nil {
		t.Error(err)
	}

	if reflect.TypeOf(uv).String() != "*openweathermap.UV" {
		t.Error("incorrect data type returned")
	}

	expected := time.Duration(1) * time.Second
	if uv.client.Timeout != expected {
		t.Errorf("Expected Duration %v, but got %v", expected, uv.client.Timeout)
	}
}

// TestNewUVWithInvalidOptions will verify that returns an error with
// invalid option
func TestNewUVWithInvalidOptions(t *testing.T) {

	optionsPattern := [][]Option{
		{nil},
		{nil, nil},
		{WithHttpClient(&http.Client{}), nil},
		{nil, WithHttpClient(&http.Client{})},
	}

	for _, options := range optionsPattern {
		c, err := NewUV(os.Getenv("OWM_API_KEY"), options...)
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

// TestNewUVWithInvalidHttpClient will verify that returns an error with
// invalid http client
func TestNewUVWithInvalidHttpClient(t *testing.T) {

	uv, err := NewUV(os.Getenv("OWM_API_KEY"), WithHttpClient(nil))
	if err == errInvalidHttpClient {
		t.Logf("Received expected bad client error. message: %s", err.Error())
	} else if err != nil {
		t.Errorf("Expected %v, but got %v", errInvalidHttpClient, err)
	}
	if uv != nil {
		t.Errorf("Expected nil, but got %v", uv)
	}
}

// TestNewUVWithApiURL  will verify that a new instance of UVWeatherData
// is created with custom API url
func TestNewUVWithApiURL(t *testing.T) {
	c, err := NewUV(os.Getenv("OWM_API_KEY"), WithApiURL("https://ru.openweathermap.org/"))
	if err != nil {
		t.Error(err)
	}

	if reflect.TypeOf(c).String() != "*openweathermap.UV" {
		t.Error("incorrect data type returned")
	}

	expected := "https://ru.openweathermap.org/data/2.5/"
	got := c.uvURL
	if got != expected {
		t.Errorf("Expected buvURL %v, but got %v", expected, got)
	}
}

// TestNewUVWithInvalidApiURL  will verify that returns an error with
// invalid API url
func TestNewUVWithInvalidApiURL(t *testing.T) {
	c, err := NewUV(os.Getenv("OWM_API_KEY"), WithApiURL("somestring"))
	if err == errInvalidApiURL {
		t.Logf("Received expected invalid url error. message: %s", err.Error())
	} else if err != nil {
		t.Errorf("Expected %v, but got %v", errInvalidApiURL, err)
	}
	if c != nil {
		t.Errorf("Expected nil, but got %v", c)
	}
}

// TestCurrentUV
func TestCurrentUV(t *testing.T) {
	t.Parallel()

	uv, err := NewUV(os.Getenv("OWM_API_KEY"))
	if err != nil {
		t.Error(err)
	}

	if err := uv.Current(coords); err != nil {
		t.Error(err)
	}

	if reflect.TypeOf(uv).String() != "*openweathermap.UV" {
		t.Error("incorrect data type returned")
	}
}

// TestHistoricalUV
func TestHistoricalUV(t *testing.T) {
	t.Parallel()

	/*	uv := NewUV(os.Getenv("OWM_API_KEY"))

		end := time.Now().UTC()
		start := time.Now().UTC().Add(-time.Hour * time.Duration(24))

		if err := uv.Historical(coords, start, end); err != nil {
			t.Error(err)
		}

		if reflect.TypeOf(uv).String() != "*openweathermap.UV" {
			t.Error("incorrect data type returned")
		}*/
}

func TestUVInformation(t *testing.T) {
	t.Parallel()

	uv, err := NewUV(os.Getenv("OWM_API_KEY"))
	if err != nil {
		t.Error(err)
	}

	if err := uv.Current(coords); err != nil {
		t.Error(err)
	}

	_, err = uv.UVInformation()
	if err != nil {
		t.Error(err)
	}
}
