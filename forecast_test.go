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

var forecastRange = []int{3, 7, 10}

// TestNewForecast will make sure the a new instance of Forecast is returned
func TestNewForecast(t *testing.T) {
	t.Parallel()

	for d := range DataUnits {
		t.Logf("Data unit: %s", d)

		if ValidDataUnit(d) {
			c, err := NewForecast(d, "ru")
			if err != nil {
				t.Error(err)
			}

			if reflect.TypeOf(c).String() != "*openweathermap.ForecastWeatherData" {
				t.Error("incorrect data type returned")
			}
		} else {
			t.Errorf("unusable data unit - %s", d)
		}
	}

	_, err := NewForecast("asdf", "en")
	if err == nil {
		t.Error("created instance when it shouldn't have")
	}
}

// TestDailyByName will verify that a daily forecast can be retrieved for
// a given named location
func TestDailyByName(t *testing.T) {
	t.Parallel()

	f, err := NewForecast("f", "fi")
	if err != nil {
		t.Error(err)
	}

	for _, d := range forecastRange {
		f.DailyByName("Dubai", d)
	}
}

// TestDailyByCooridinates will verify that a daily forecast can be retrieved
// for a given set of coordinates
func TestDailyByCoordinates(t *testing.T) {
	t.Parallel()

	f, err := NewForecast("f", "PL")
	if err != nil {
		t.Error(err)
	}

	for _, d := range forecastRange {
		f.DailyByCoordinates(
			&Coordinates{
				Longitude: -112.07,
				Latitude:  33.45,
			}, d,
		)
	}
}

// TestDailyByID will verify that a daily forecast can be retrieved for a
// given location ID
func TestDailyByID(t *testing.T) {
	t.Parallel()

	f, err := NewForecast("c", "fr")
	if err != nil {
		t.Error(err)
	}

	for _, d := range forecastRange {
		f.DailyByID(524901, d)
	}
}
