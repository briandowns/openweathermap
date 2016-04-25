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
	"reflect"
	"testing"
)

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

type current_weather struct {
	query		string
	weather		CurrentWeatherData
}
// TestCurrentByName will verify that current data can be retrieved for a given
// location by name
func TestCurrentByName(t *testing.T) {
	t.Parallel()
	testCities := []current_weather {
		{
			query:		"Philadelphia",
			weather:	CurrentWeatherData {
				ID:		4560349,
				Name:	"Philadelphia",
				Main: Main {
					Temp:	35.6,
				},
			},
		},
		{
			query:		"Newark",
			weather:	CurrentWeatherData {
				ID:		5101798,
				Name:	"Newark",
				Main: Main {
					Temp:	36.36,
				},
			},
		},
		{
			query:		"Helena",
			weather:	CurrentWeatherData {
				ID:		5656882,
				Name:	"Helena",
				Main: Main {
					Temp:	42.8,
				},
			},
		},
		{
			query:		"San Diego, CA",
			weather:	CurrentWeatherData {
				ID:		5391811,
				Name:	"San Diego",
				Main: Main {
					Temp:	56.53,
				},
			},
		},
	}

	c, err := NewCurrent("f", "ru")
	if err != nil {
		t.Error(err)
	}
	for _, city := range testCities {
		c.CurrentByName(city.query)

		if c.ID != city.weather.ID {
			t.Errorf("Excpect CityID %ld, got %ld", city.weather.ID, c.ID)
		}
		if c.Name != city.weather.Name {
			t.Errorf("Excpect City %s, got %s", city.weather.Name, c.Name)
		}
		if c.Main.Temp != city.weather.Main.Temp {
			t.Errorf("Excpect Temp %.2f, got %.2f", city.weather.Main.Temp, c.Main.Temp)
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
