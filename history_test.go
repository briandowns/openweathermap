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

// TestNewHistory verifies NewHistorical does as advertised
func TestNewHistory(t *testing.T) {
	t.Parallel()

	for d := range DataUnits {
		t.Logf("Data unit: %s", d)

		if ValidDataUnit(d) {
			c, err := NewHistorical(d, os.Getenv("OWM_API_KEY"))
			if err != nil {
				t.Error(err)
			}
			if reflect.TypeOf(c).String() != "*openweathermap.HistoricalWeatherData" {
				t.Error("incorrect data type returned")
			}
		} else {
			t.Errorf("unusable data unit - %s", d)
		}
	}

	_, err := NewHistorical("asdf", os.Getenv("OWM_API_KEY"))
	if err == nil {
		t.Error("created instance when it shouldn't have")
	}
}

// TestNewHistoryWithCustomHttpClient will verify that a new instance of HistoricalWeatherData
// is created with custom http client
func TestNewHistoryWithCustomHttpClient(t *testing.T) {

	hc := http.DefaultClient
	hc.Timeout = time.Duration(1) * time.Second
	h, err := NewHistorical("c", os.Getenv("OWM_API_KEY"), WithHttpClient(hc))
	if err != nil {
		t.Error(err)
	}

	if reflect.TypeOf(h).String() != "*openweathermap.HistoricalWeatherData" {
		t.Error("incorrect data type returned")
	}

	expected := time.Duration(1) * time.Second
	if h.client.Timeout != expected {
		t.Errorf("Expected Duration %v, but got %v", expected, h.client.Timeout)
	}
}

// TestNewHistoryWithInvalidOptions will verify that returns an error with
// invalid option
func TestNewHistoryWithInvalidOptions(t *testing.T) {

	optionsPattern := [][]Option{
		{nil},
		{nil, nil},
		{WithHttpClient(&http.Client{}), nil},
		{nil, WithHttpClient(&http.Client{})},
	}

	for _, options := range optionsPattern {
		c, err := NewHistorical("c", os.Getenv("OWM_API_KEY"), options...)
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

// TestNewHistoryWithInvalidHttpClient will verify that returns an error with
// invalid http client
func TestNewHistoryWithInvalidHttpClient(t *testing.T) {

	h, err := NewHistorical("c", os.Getenv("OWM_API_KEY"), WithHttpClient(nil))
	if err == errInvalidHttpClient {
		t.Logf("Received expected bad client error. message: %s", err.Error())
	} else if err != nil {
		t.Errorf("Expected %v, but got %v", errInvalidHttpClient, err)
	}
	if h != nil {
		t.Errorf("Expected nil, but got %v", h)
	}
}

// TestNewHistorylWithApiURL  will verify that a new instance of HisoricalWeatherData
// is created with custom API url
func TestNewHistoryWithApiURL(t *testing.T) {
	c, err := NewHistorical("c", os.Getenv("OWM_API_KEY"), WithApiURL("https://ru.openweathermap.org/"))
	if err != nil {
		t.Error(err)
	}

	if reflect.TypeOf(c).String() != "*openweathermap.HistoricalWeatherData" {
		t.Error("incorrect data type returned")
	}

	expected := "https://ru.openweathermap.org/data/2.5/history/%s"
	got := c.historyURL
	if got != expected {
		t.Errorf("Expected historyURL %v, but got %v", expected, got)
	}
}

// TestNewHistoryWithInvalidApiURL  will verify that returns an error with
// invalid API url
func TestNewHistoryWithInvalidApiURL(t *testing.T) {
	c, err := NewHistorical("c", os.Getenv("OWM_API_KEY"), WithApiURL("somestring"))
	if err == errInvalidApiURL {
		t.Logf("Received expected invalid url error. message: %s", err.Error())
	} else if err != nil {
		t.Errorf("Expected %v, but got %v", errInvalidApiURL, err)
	}
	if c != nil {
		t.Errorf("Expected nil, but got %v", c)
	}
}

// TestHistoryByName
func TestHistoryByName(t *testing.T) {
	t.Parallel()
	h, err := NewHistorical("F", os.Getenv("OWM_API_KEY"))
	if err != nil {
		t.Error(err)
	}
	if err := h.HistoryByName("Vancouver"); err != nil {
		t.Error(err)
	}
}

// TestHistoryByID
func TestHistoryByID(t *testing.T) {
	t.Parallel()
	h, err := NewHistorical("F", os.Getenv("OWM_API_KEY"))
	if err != nil {
		t.Error(err)
	}
	hp := &HistoricalParameters{
		Start: 1461598510,
		End:   1461588510,
		Cnt:   1,
	}
	if err := h.HistoryByID(5344157, hp); err != nil {
		t.Error(err)
	}
}

// TestHistoryByCoord
func TestHistoryByCoord(t *testing.T) {
	t.Parallel()
	h, err := NewHistorical("F", os.Getenv("OWM_API_KEY"))
	if err != nil {
		t.Error(err)
	}
	coords := &Coordinates{
		Longitude: -112.07,
		Latitude:  33.45,
	}
	hp := &HistoricalParameters{
		Start: 1461598510,
		End:   1461588510,
		Cnt:   1,
	}
	if err := h.HistoryByCoord(coords, hp); err != nil {
		t.Error(err)
	}
}
