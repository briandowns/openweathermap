package openweathermap

import (
	"os"
	"testing"
)

// TestCurrentByIDs will verify that list of current data can be retrieved for
// a given by location ids
func TestCurrentByIDs(t *testing.T) {
	t.Parallel()
	g, err := NewCurrentGroup("c", "RU", os.Getenv("OWM_API_KEY"))
	if err != nil {
		t.Error("Error creating instance of CurrentWeatherData")
	}

	// testing error
	ids := make([]int, 90)
	for i := 0; i < 90; i++ {
		ids[i] = i + 1
	}
	err = g.CurrentByIDs(ids...)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	// test receiving real data
	err = g.CurrentByIDs(5344157, 5391811, 4560349)
	if err != nil {
		t.Fatalf("getting list by ids of current waeather failed: %s", err)
	}

	if n := len(g.List); n != 3 {
		t.Errorf("wrong count of results: expected %d, got %d", 3, n)
	}
}
