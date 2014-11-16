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

func TestNewForecast(t *testing.T) {
	t.Parallel()
	for _, u := range dataUnits {
		t.Logf("Data unit: %s", u)
		if ValidDataUnit(u) {
			c, err := NewForecast(u)
			if err != nil {
				t.Error(err)
			}
			if reflect.TypeOf(c).String() != "*openweathermap.ForecastWeatherData" {
				t.Error("ERROR: incorrect data type returned")
			}
		} else {
			t.Errorf("ERROR: unusable data unit - %s", u)
		}
	}
}

func TestDailyByName(t *testing.T) {
	f, err := NewForecast("imperial")
	if err != nil {
		t.Error("")
	}
	f.DailyByName("Dubai", 3)
}

func TestDailyByCoordinates(t *testing.T) {}
func TestDailyByID(t *testing.T)          {}
