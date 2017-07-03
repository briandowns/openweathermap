// Copyright 2017 Gwennaël Buchet
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
	"testing"
)

var hourlyForecastRange = []int{3, 7, 10}

// TestForecastByName will verify that a daily forecast can be retrieved for
// a given named location
func TestHourlyForecastByName(t *testing.T) {
	t.Parallel()

	f, err := NewHourlyForecast("c", "fr")
	if err != nil {
		t.Error(err)
	}

	for _, d := range hourlyForecastRange {
		f.HourlyForecastByName("Lille", d)
	}
}

// TestForecastByCooridinates will verify that a daily forecast can be retrieved
// for a given set of coordinates
func TestHourlyForecastByCoordinates(t *testing.T) {
	t.Parallel()

	f, err := NewHourlyForecast("f", "PL")
	if err != nil {
		t.Error(err)
	}

	for _, d := range hourlyForecastRange {
		f.HourlyForecastByCoordinates(
			&Coordinates{
				Longitude: -112.07,
				Latitude:  33.45,
			}, d,
		)
	}
}

// TestForecastByID will verify that a daily forecast can be retrieved for a
// given location ID
func TestHourlyForecastByID(t *testing.T) {
	t.Parallel()

	f, err := NewHourlyForecast("c", "fr")
	if err != nil {
		t.Error(err)
	}

	for _, d := range hourlyForecastRange {
		f.HourlyForecastByID(524901, d)
	}
}
