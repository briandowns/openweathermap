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
	owm "github.com/briandowns/openweathermap"
	"reflect"
	"testing"
)

// TestNewHistory verifies NewHistorical does as advertised
func TestNewHistory(t *testing.T) {
	t.Parallel()
	for d, _ := range owm.DataUnits {
		t.Logf("Data unit: %s", d)
		if owm.ValidDataUnit(d) {
			c, err := owm.NewHistorical(d)
			if err != nil {
				t.Error(err)
			}
			if reflect.TypeOf(c).String() != "*openweathermap.HistoricalWeatherData" {
				t.Error("ERROR: incorrect data type returned")
			}
		} else {
			t.Errorf("ERROR: unusable data unit - %s", d)
		}
	}
	_, err := owm.NewHistorical("asdf")
	if err == nil {
		t.Error("ERROR: created instance when it shouldn't have")
	}
}
