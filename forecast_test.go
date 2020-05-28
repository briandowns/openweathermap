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
	"os"
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
			c5, err := NewForecast("5", d, "ru", os.Getenv("OWM_API_KEY"))
			if err != nil {
				t.Error(err)
			}

			if reflect.TypeOf(c5).String() != "*openweathermap.ForecastWeatherData" {
				t.Error("incorrect data type returned")
			}

			c16, err := NewForecast("16", d, "ru", os.Getenv("OWM_API_KEY"))
			if err != nil {
				t.Error(err)
			}

			if reflect.TypeOf(c16).String() != "*openweathermap.ForecastWeatherData" {
				t.Error("incorrect data type returned")
			}
		} else {
			t.Errorf("unusable data unit - %s", d)
		}
	}

	_, err := NewForecast("", "asdf", "en", os.Getenv("OWM_API_KEY"))
	if err == nil {
		t.Error("created instance when it shouldn't have")
	}
}

// TestNewForecastWithCustomHttpClient will verify that a new instance of ForecastWeatherData
// is created with custom http client
func TestNewForecastWithCustomHttpClient(t *testing.T) {

	hc := http.DefaultClient
	hc.Timeout = time.Duration(1) * time.Second
	f, err := NewForecast("5", "c", "en", os.Getenv("OWM_API_KEY"), WithHttpClient(hc))
	if err != nil {
		t.Error(err)
	}

	if reflect.TypeOf(f).String() != "*openweathermap.ForecastWeatherData" {
		t.Error("incorrect data type returned")
	}

	// expected := time.Duration(1) * time.Second
	// if f.client.Timeout != expected {
	// 	t.Errorf("Expected Duration %v, but got %v", expected, f.client.Timeout)
	// }
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
		c, err := NewForecast("5", "c", "en", os.Getenv("OWM_API_KEY"), options...)
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

	f, err := NewForecast("5", "c", "en", os.Getenv("OWM_API_KEY"), WithHttpClient(nil))
	if err == errInvalidHttpClient {
		t.Logf("Received expected bad client error. message: %s", err.Error())
	} else if err != nil {
		t.Errorf("Expected %v, but got %v", errInvalidHttpClient, err)
	}
	if f != nil {
		t.Errorf("Expected nil, but got %v", f)
	}
}

// TestNewForecast5WithApiURL  will verify that a new instance of ForecastWeatherData
// is created with custom API url
func TestNewForecast5WithApiURL(t *testing.T) {
	c, err := NewForecast("5", "c", "en", os.Getenv("OWM_API_KEY"), WithApiURL("https://ru.openweathermap.org/"))
	if err != nil {
		t.Error(err)
	}

	if reflect.TypeOf(c).String() != "*openweathermap.ForecastWeatherData" {
		t.Error("incorrect data type returned")
	}

	expected := "https://ru.openweathermap.org/data/2.5/forecast?appid=%s&%s&mode=json&units=%s&lang=%s&cnt=%d"
	got := c.baseURL
	if got != expected {
		t.Errorf("Expected baseURL %v, but got %v", expected, got)
	}
}

// TestNewForecast16WithApiURL  will verify that a new instance of ForecastWeatherData
// is created with custom API url
func TestNewForecast16WithApiURL(t *testing.T) {
	c, err := NewForecast("16", "c", "en", os.Getenv("OWM_API_KEY"), WithApiURL("https://ru.openweathermap.org/"))
	if err != nil {
		t.Error(err)
	}

	if reflect.TypeOf(c).String() != "*openweathermap.ForecastWeatherData" {
		t.Error("incorrect data type returned")
	}

	expected := "https://ru.openweathermap.org/data/2.5/forecast/daily?appid=%s&%s&mode=json&units=%s&lang=%s&cnt=%d"
	got := c.baseURL
	if got != expected {
		t.Errorf("Expected baseURL %v, but got %v", expected, got)
	}
}

// TestNewForecastWithInvalidApiURL  will verify that returns an error with
// invalid API url
func TestNewForecastWithInvalidApiURL(t *testing.T) {
	c, err := NewForecast("5", "c", "en", os.Getenv("OWM_API_KEY"), WithApiURL("somestring"))
	if err == errInvalidApiURL {
		t.Logf("Received expected invalid url error. message: %s", err.Error())
	} else if err != nil {
		t.Errorf("Expected %v, but got %v", errInvalidApiURL, err)
	}
	if c != nil {
		t.Errorf("Expected nil, but got %v", c)
	}
}

// TestDailyByName will verify that a daily forecast can be retrieved for
// a given named location
func TestDailyByName(t *testing.T) {
	t.Parallel()

	f, err := NewForecast("5", "f", "fi", os.Getenv("OWM_API_KEY"))
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

	f, err := NewForecast("5", "f", "PL", os.Getenv("OWM_API_KEY"))
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

	f, err := NewForecast("5", "c", "fr", os.Getenv("OWM_API_KEY"))
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
