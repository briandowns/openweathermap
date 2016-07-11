package openweathermap

import (
	"net/http"
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

	uv, err := NewUV()
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
	uv, err := NewUV(WithHttpClient(hc))
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

// TestCurrentUV
func TestCurrentUV(t *testing.T) {
	t.Parallel()

	uv, err := NewUV()
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

	/*	uv := NewUV()

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

	uv, err := NewUV()
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
