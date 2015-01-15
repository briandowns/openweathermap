// Copyright 2014 Brian J. Downs
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package openweathermap

import (
	"reflect"
	"testing"
)

// TestNewCurrent will verify that a new instance of CurrentWeatherData is created
func TestNewCurrent(t *testing.T) {
	t.Parallel()
	for d, _ := range DataUnits {
		t.Logf("Data unit: %s", d)
		if ValidDataUnit(d) {
			c, err := NewCurrent(d, "en")
			if err != nil {
				t.Error(err)
			}
			if reflect.TypeOf(c).String() != "*openweathermap.CurrentWeatherData" {
				t.Error("incorrect data type returned")
			}
		} else {
			t.Errorf("unusable data unit - %s", d)
		}
	}
	_, err := NewCurrent("Philadelphia", "en")
	if err == nil {
		t.Error("created instance when it shouldn't have")
	}
}

// TestCurrentByName will verify that current data can be retrieved for a give
// location by name
func TestCurrentByName(t *testing.T) {
	t.Parallel()
	testCities := []string{"Philadelphia", "Newark", "Helena"}
	c, err := NewCurrent("f", "ru")
	if err != nil {
		t.Error(err)
	}
	for _, city := range testCities {
		c.CurrentByName(city)
		if c.Name != city {
			t.Error("incorrect city returned")
		}
	}
}

// TestCurrentByCoordinates will verify that current data can be retrieved for a
// given set of coordinates
func TestCurrentByCoordinates(t *testing.T) {
	t.Parallel()
	c, err := NewCurrent("f", "DE")
	if err != nil {
		t.Error("Error creating instance of CurrentWeatherData")
	}
	c.CurrentByCoordinates(
		&Coordinates{
			Longitude: -112.07,
			Latitude:  33.45,
		},
	)
}

// TestCurrentByID will verify that current data can be retrieved for a given
// location id
func TestCurrentByID(t *testing.T) {
	t.Parallel()
	c, err := NewCurrent("c", "ZH")
	if err != nil {
		t.Error("Error creating instance of CurrentWeatherData")
	}
	c.CurrentByID(5344157)
}

func TestCurrentByArea(t *testing.T) {}
