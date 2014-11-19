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

var forecastRange = []int{3, 7, 10}

func TestNewForecast(t *testing.T) {
	t.Parallel()
	for d, _ := range dataUnits {
		t.Logf("Data unit: %s", d)
		if ValidDataUnit(d) {
			c, err := NewForecast(d)
			if err != nil {
				t.Error(err)
			}
			if reflect.TypeOf(c).String() != "*openweathermap.ForecastWeatherData" {
				t.Error("ERROR: incorrect data type returned")
			}
		} else {
			t.Errorf("ERROR: unusable data unit - %s", d)
		}
	}
	_, err := NewForecast("asdf")
	if err == nil {
		t.Error("ERROR: created instance when it shouldn't have")
	}
}

func TestDailyByName(t *testing.T) {
	f, err := NewForecast("imperial")
	if err != nil {
		t.Error(err)
	}
	for _, d := range forecastRange {
		f.DailyByName("Dubai", d)
	}
}

func TestDailyByCoordinates(t *testing.T) {
	f, err := NewForecast("internal")
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

func TestDailyByID(t *testing.T) {
	f, err := NewForecast("metric")
	if err != nil {
		t.Error(err)
	}
	for _, d := range forecastRange {
		f.DailyByID(524901, d)
	}
}
