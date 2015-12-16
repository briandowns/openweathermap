package openweathermap

import "testing"

// TestCurrentUV
func TestCurrentUV(t *testing.T) {
	t.Parallel()

	/*coord := &Coordinates{
		Longitude: 53.343497,
		Latitude:  -6.288379,
	}

	c, err := CurrentUV(coord)
	if err != nil {
		t.Error(err)
	}

	if reflect.TypeOf(c).String() != "*openweathermap.CurUV" {
		t.Error("incorrect data type returned")
	}*/
}

// TestHistoricalUV
func TestHistoricalUV(t *testing.T) {
	t.Parallel()

	/*// 54.995656, -7.326834
	coord := &Coordinates{
		Longitude: 54.995656,
		Latitude:  -7.326834,
	}

	end := time.Now().UTC()
	start := time.Now().UTC().Add(-time.Hour * time.Duration(24))

	h, err := HistoricalUV(coord, start, end)
	if err != nil {
		t.Error(err)
	}

	if reflect.TypeOf(h).String() != "*openweathermap.CurUV" {
		t.Error("incorrect data type returned")
	}*/
}
