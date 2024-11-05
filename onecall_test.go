// Copyright 2022 Brian J. Downs
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

// TestNewOneCall will verify that a new instance of OneCallData is created
func TestNewOneCall(t *testing.T) {
	t.Parallel()

	for d := range DataUnits {
		t.Logf("Data unit: %s", d)

		if ValidDataUnit(d) {
			c, err := NewOneCall(d, "en", os.Getenv("OWM_API_KEY"), []string{})
			if err != nil {
				t.Error(err)
			}

			if _, err := NewOneCall(d, "blah", os.Getenv("OWM_API_KEY"), []string{}); err != nil {
				t.Log("received expected bad language code error")
			}

			if reflect.TypeOf(c).String() != "*openweathermap.OneCallData" {
				t.Error("incorrect data type returned")
			}
		} else {
			t.Errorf("unusable data unit - %s", d)
		}
	}
}

// TestNewOneCallWithCustomHttpClient will verify that a new instance of OneCallData
// is created with custom http client
func TestNewOneCallWithCustomHttpClient(t *testing.T) {
	hc := http.DefaultClient
	hc.Timeout = time.Duration(1) * time.Second
	c, err := NewOneCall("c", "en", os.Getenv("OWM_API_KEY"), []string{}, WithHttpClient(hc))
	if err != nil {
		t.Error(err)
	}

	if reflect.TypeOf(c).String() != "*openweathermap.OneCallData" {
		t.Error("incorrect data type returned")
	}

	expected := time.Duration(1) * time.Second
	if c.client.Timeout != expected {
		t.Errorf("Expected Duration %v, but got %v", expected, c.client.Timeout)
	}
}

// TestNewOneCallWithInvalidOptions will verify that returns an error with
// invalid option
func TestNewOneCallWithInvalidOptions(t *testing.T) {
	optionsPattern := [][]Option{
		{nil},
		{nil, nil},
		{WithHttpClient(&http.Client{}), nil},
		{nil, WithHttpClient(&http.Client{})},
	}

	for _, options := range optionsPattern {
		c, err := NewOneCall("c", "en", os.Getenv("OWM_API_KEY"), []string{}, options...)
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

// TestNewCurrentWithInvalidHttpClient will verify that returns an error with
// invalid http client
func TestNewOneCallWithInvalidHttpClient(t *testing.T) {

	c, err := NewOneCall("c", "en", os.Getenv("OWM_API_KEY"), []string{}, WithHttpClient(nil))
	if err == errInvalidHttpClient {
		t.Logf("Received expected bad client error. message: %s", err.Error())
	} else if err != nil {
		t.Errorf("Expected %v, but got %v", errInvalidHttpClient, err)
	}
	if c != nil {
		t.Errorf("Expected nil, but got %v", c)
	}
}

// TestOneCallByCoordinates will verify that onecall data can be retrieved for a
// given set of coordinates
func TestOneCallByCoordinates(t *testing.T) {
	t.Parallel()
	c, err := NewOneCall("f", "DE", os.Getenv("OWM_API_KEY"), []string{})
	if err != nil {
		t.Error("Error creating instance of OneCallData")
	}
	err = c.OneCallByCoordinates(
		&Coordinates{
			Longitude: -112.07,
			Latitude:  33.45,
		},
	)
	if err != nil {
		t.Error(err)
	}
}

func TestNewOneCallWithOneExclude(t *testing.T) {
	c, err := NewOneCall("f", "en", os.Getenv("OWM_API_KEY"), []string{ExcludeAlerts})
	if err != nil {
		t.Error(err)
	}

	err = c.OneCallByCoordinates(
		&Coordinates{
			Longitude: -112.07,
			Latitude:  33.45,
		},
	)
	if err != nil {
		t.Error(err)
	}

	if len(c.Alerts) > 0 {
		t.Error("exclude alerts fails")
	}
}

func TestNewOneCallWithTwoExcludes(t *testing.T) {
	c, err := NewOneCall("f", "en", os.Getenv("OWM_API_KEY"), []string{ExcludeAlerts, ExcludeDaily})
	if err != nil {
		t.Error(err)
	}

	err = c.OneCallByCoordinates(
		&Coordinates{
			Longitude: -112.07,
			Latitude:  33.45,
		},
	)
	if err != nil {
		t.Error(err)
	}

	if len(c.Alerts) > 0 && len(c.Daily) > 0 {
		t.Error("exclude alerts and daily fails")
	}
}

// TestOneCallTimeMachine will verify that onecall data can be retrieved for a
// given set of coordinates and a time
func TestOneCallTimeMachine(t *testing.T) {
	t.Parallel()
	c, err := NewOneCall("f", "DE", os.Getenv("OWM_API_KEY"), []string{})
	if err != nil {
		t.Error("Error creating instance of OneCallData")
	}
	err = c.OneCallTimeMachine(
		&Coordinates{
			Longitude: -112.07,
			Latitude:  33.45,
		},
		time.Now().AddDate(0, 0, -1),
	)
	if err != nil {
		t.Error(err)
	}
}
