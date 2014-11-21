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
	"testing"
)

// TestValidDataUnit tests whether or not ValidDataUnit provides
// the correct assertion on provided data unit.
func TestValidDataUnit(t *testing.T) {
	for u, _ := range owm.DataUnits {
		if !owm.ValidDataUnit(u) {
			t.Error("False positive on data unit")
		}
	}
	if owm.ValidDataUnit("anything") {
		t.Error("Invalid data unit")
	}
}

func TestDataUnitValues(t *testing.T) {
	for _, s := range owm.DataUnits {
		if !owm.ValidDataUnitSymbol(s) {
			t.Error("False positive on data unit symbol")
		}
	}
	if owm.ValidDataUnitSymbol("X") {
		t.Error("Invalid data unit symbol")
	}
}
