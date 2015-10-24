// Copyright 2015 Brian J. Downs
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
	"os"
	"reflect"
	"testing"
)

func testSetup() {
	Config.SetApiKey(os.Getenv("OWM_API_KEY"))
}

func testTeardown() {
	Config.SetApiKey("")
}

// TestNewCurrent will verify that a new instance of CurrentWeatherData is created
func TestNewCurrent(t *testing.T) {
	t.Parallel()
	for d := range DataUnits {
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
	testSetup()
	testCities := []string{"Philadelphia", "Newark", "Helena", "San Diego, CA"}
	c, err := NewCurrent("f", "ru")
	if err != nil {
		t.Error(err)
	}
	for _, city := range testCities {
		if bFetchErr := c.CurrentByName(city); bFetchErr != nil {
			t.Error(bFetchErr)
		}
	}
	testTeardown()
}

// TestCurrentByCoordinates will verify that current data can be retrieved for a
// given set of coordinates
func TestCurrentByCoordinates(t *testing.T) {
	testSetup()
	c, err := NewCurrent("f", "DE")
	if err != nil {
		t.Error("Error creating instance of CurrentWeatherData")
	}
	bFetchErr := c.CurrentByCoordinates(
		&Coordinates{
			Longitude: -112.07,
			Latitude:  33.45,
		},
	)

	if bFetchErr != nil {
		t.Error(bFetchErr)
	}
	testTeardown()
}

// TestCurrentByID will verify that current data can be retrieved for a given
// location id
func TestCurrentByID(t *testing.T) {
	testSetup()
	c, err := NewCurrent("c", "ZH")
	if err != nil {
		t.Error("Error creating instance of CurrentWeatherData")
	}
	bFetchErr := c.CurrentByID(5344157)

	if bFetchErr != nil {
		t.Error(bFetchErr)
	}
	testTeardown()
}

func TestCurrentByArea(t *testing.T) {}
