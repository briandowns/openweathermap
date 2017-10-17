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
	"net/http"
	"reflect"
	"testing"
	"time"
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

// TestNewForecastWithCustomHttpClient will verify that a new instance of ForecastWeatherData
// is created with custom http client
func TestNewForecastWithCustomHttpClient(t *testing.T) {

	hc := http.DefaultClient
	hc.Timeout = time.Duration(1) * time.Second
	f, err := NewForecast("c", "en", WithHttpClient(hc))
	if err != nil {
		t.Error(err)
	}

	if reflect.TypeOf(f).String() != "*openweathermap.ForecastWeatherData" {
		t.Error("incorrect data type returned")
	}

	expected := time.Duration(1) * time.Second
	if f.client.Timeout != expected {
		t.Errorf("Expected Duration %v, but got %v", expected, f.client.Timeout)
	}
}

// TestNewForecastWithInvalidOptions will verify that returns an error with
// invalid option
func TestNewForecastWithInvalidOptions(t *testing.T) {

	optionsPattern := [][]Option{
		{nil},
		{nil, nil},
		{WithHttpClient(&http.Client{}), nil},
		{nil, WithHttpClient(&http.Client{})},
	}

	for _, options := range optionsPattern {
		c, err := NewForecast("c", "en", options...)
		if err == errInvalidOption {
			t.Logf("Received expected invalid option error. message: %s", err.Error())
		} else if err != nil {
			t.Errorf("Expected %v, but got %v", errInvalidOption, err)
		}
		if c != nil {
			t.Errorf("Expected nil, but got %v", c)
		}
	}
}

// TestNewForecastWithCustomHttpClient will verify that returns an error with
// invalid http client
func TestNewForecastWithInvalidHttpClient(t *testing.T) {

	f, err := NewForecast("c", "en", WithHttpClient(nil))
	if err == errInvalidHttpClient {
		t.Logf("Received expected bad client error. message: %s", err.Error())
	} else if err != nil {
		t.Errorf("Expected %v, but got %v", errInvalidHttpClient, err)
	}
	if f != nil {
		t.Errorf("Expected nil, but got %v", f)
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
		err = f.DailyByName("Dubai", d)
		if err != nil {
			t.Error(err)
		}
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
		err = f.DailyByCoordinates(
			&Coordinates{
				Longitude: -112.07,
				Latitude:  33.45,
			}, d,
		)
		if err != nil {
			t.Error(err)
		}
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
		err = f.DailyByID(524901, d)
		if err != nil {
			t.Error(err)
		}
	}
}
