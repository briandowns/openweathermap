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
	"testing"
)

// TestValidDataUnit tests whether or not ValidDataUnit provides
// the correct assertion on provided data unit.
func TestValidDataUnit(t *testing.T) {
	for u := range DataUnits {
		if !ValidDataUnit(u) {
			t.Error("False positive on data unit")
		}
	}

	if ValidDataUnit("anything") {
		t.Error("Invalid data unit")
	}
}

func TestDataUnitValues(t *testing.T) {
	for _, s := range DataUnits {
		if !ValidDataUnitSymbol(s) {
			t.Error("False positive on data unit symbol")
		}
	}

	if ValidDataUnitSymbol("X") {
		t.Error("Invalid data unit symbol")
	}
}

func TestCheckAPIKeyExists(t *testing.T) {
	c := &Config{
		APIKey: "asdf1234",
	}

	if !c.CheckAPIKeyExists() {
		t.Error("Key not set")
	}
}

// TestSetOptionsWithEmpty tests setOptions function will do nothing
// when options are empty.
func TestSetOptionsWithEmpty(t *testing.T) {
	s := NewSettings()
	err := setOptions(s, nil)
	if err != nil {
		t.Error(err)
	}
}

// TestValidAPIKey tests ValidAPIKey function can correctly identify properly
// and improperly formatted API keys.
func TestValidAPIKey(t *testing.T) {
	casesOk := []struct {
		key string
	}{
		{
			key: "36422f06455843a1b4131b475f2c10c0",
		},
		{
			key: "587119ac943749d69037308fe4e810d7",
		},
		{
			key: "ea6a902532f24e8fa89c5fc8fabb6f98",
		},
	}
	casesNotOk := []struct {
		key string
		err error
	}{
		{
			key: "98291d5835f442688226x6z79eadcf7",
			err: errInvalidKey,
		},
		{
			key: "1924akd992j4342be96d82q46c38ea8",
			err: errInvalidKey,
		},
		{
			key: "ea6a902532l24p8fa89c5fc8fabb6f98",
			err: errInvalidKey,
		},
	}
	for _, v := range casesOk {
		if err := ValidAPIKey(v.key); err != nil {
			t.Errorf("Key: %s, %v", v.key, err)
		}
	}
	for _, v := range casesNotOk {
		if err := ValidAPIKey(v.key); err != nil {
			if err != v.err {
				t.Errorf("Expected %v error; got %v instead", v.err, err)
			}
		} else {
			t.Errorf("Expected %v error; got nil instead", v.err)
		}
	}
}
