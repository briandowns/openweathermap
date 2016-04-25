package openweathermap

import (
	"reflect"
	"testing"
)

var coords = &Coordinates{
	Longitude: 53.343497,
	Latitude:  -6.288379,
}

// TestCurrentUV
func TestCurrentUV(t *testing.T) {
	t.Parallel()

	uv := NewUV()

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

	uv := NewUV()

	if err := uv.Current(coords); err != nil {
		t.Error(err)
	}

	_, err := uv.UVInformation()
	if err != nil {
		t.Error(err)
	}
}
