package openweathermap

import (
	"reflect"
	"testing"
)

// TestCurrentUV
func TestCurrentUV(t *testing.T) {
	t.Parallel()

	uv := NewUV()

	coord := &Coordinates{
		Longitude: 53.343497,
		Latitude:  -6.288379,
	}

	if err := uv.Current(coord); err != nil {
		t.Error(err)
	}

	if reflect.TypeOf(uv).String() != "*openweathermap.UV" {
		t.Error("incorrect data type returned")
	}
}

// TestHistoricalUV
func TestHistoricalUV(t *testing.T) {
	t.Parallel()

	/*uv := NewUV()

	coord := &Coordinates{
		Longitude: 54.995656,
		Latitude:  -7.326834,
	}

	end := time.Now().UTC()
	start := time.Now().UTC().Add(-time.Hour * time.Duration(24))

	if err := uv.Historical(coord, start, end); err != nil {
		t.Error(err)
	}

	if reflect.TypeOf(uv).String() != "*openweathermap.UV" {
		t.Error("incorrect data type returned")
	}*/
}

func TestUVInformation(t *testing.T) {
	t.Parallel()

	uvc := NewUV()

	if err := uv.Current(coord); err != nil {
		t.Error(err)
	}

	info, err := uvc.UVInformation()
	if err != nil {
		t.Error(err)
	}

	t.Log(info)
}
