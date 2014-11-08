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
	"testing"
)

func validDataUnit(h *HistoricalWeatherData) bool {
	for _, m := range dataUnits {
		if h.Units == m {
			return true
		}
	}
	return false
}

func TestNewHistorical(t *testing.T) {
	h, err := NewHistorical("imperial")
	if err != nil {
		t.Error("Failed creating instance of HistoricalWeatherData")
	}
	if !validDataUnit(h) {
		t.Error("Invalid data unit")
	}
}
