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

var forecast3HRange = []int{1, 3, 5}

// TestNewForecast3HData will make sure the a new instance of Forecast is returned
func TestNewForecast3HData(t *testing.T) {
	t.Parallel()
	for d := range DataUnits {
		t.Logf("Data unit: %s", d)
		if ValidDataUnit(d) {
			c, err := NewForecast3H(d, "ru")
			if err != nil {
				t.Error(err)
			}
			if reflect.TypeOf(c).String() != "*openweathermap.Forecast3HData" {
				t.Error("incorrect data type returned")
			}
		} else {
			t.Errorf("unusable data unit - %s", d)
		}
	}
	_, err := NewForecast3H("asdf", "en")
	if err == nil {
		t.Error("created instance when it shouldn't have")
	}
}

// TestForecast3HByName will verify that a hourly forecast can be retrieved for
// a given named location
func TestForecast3HByName(t *testing.T) {
	testSetup()
	f, err := NewForecast3H("f", "fi")
	if err != nil {
		t.Error(err)
	}
	for _, d := range forecast3HRange {
		if bFetchErr := f.ByName("Malmö", d); bFetchErr != nil {
			t.Error(bFetchErr)
		}
	}
	testTeardown()
}

// TestForecast3HByCoordinates will verify that a hourly forecast can be retrieved
// for a given set of coordinates
func TestForecast3HByCoordinates(t *testing.T) {
	testSetup()
	f, err := NewForecast3H("f", "PL")
	if err != nil {
		t.Error(err)
	}
	for _, d := range forecast3HRange {
		bFetchErr := f.ByCoordinates(
			&Coordinates{
				Longitude: -112.07,
				Latitude:  33.45,
			}, d,
		)

		if bFetchErr != nil {
			t.Error(bFetchErr)
		}
	}
	testTeardown()
}

// TestForecast3HByID will verify that a hourly forecast can be retrieved for a
// given location ID
func TestForecast3HByID(t *testing.T) {
	testSetup()
	f, err := NewForecast3H("c", "fr")
	if err != nil {
		t.Error(err)
	}
	for _, d := range forecast3HRange {
		if bFetchErr := f.ByID(524901, d); bFetchErr != nil {
			t.Error(bFetchErr)
		}
	}
	testTeardown()
}
