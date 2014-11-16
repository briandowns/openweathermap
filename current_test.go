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

func TestNewCurrent(t *testing.T) {
	for _, u := range dataUnits {
		if ValidDataUnit(u) {
			c, err := NewCurrent(u)
			if err != nil {
				t.Error(err)
			}
			if reflect.TypeOf(c).String() != "*openweathermap.CurrentWeatherData" {
				t.Error("ERROR: incorrect data type returned")
			}
		} else {
			t.Errorf("ERROR: unusable data unit - %s", u)
		}
	}
}

func TestCurrentByName(t *testing.T) {
	c, err := NewCurrent("imperial")
	if err != nil {
		t.Error("")
	}
	c.CurrentByName("Philadelphia")
}

func TestCurrentByCoordinates(t *testing.T) {

}

func TestCurrentByID(t *testing.T) {

}

func TestCurrentByArea(t *testing.T) {}
