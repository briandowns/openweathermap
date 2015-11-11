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
	"reflect"
	"testing"
)

// TestValidateStationDataParameter will make sure that a parameter passed
// to ValidateStationDataParameter is in fact a valid parameter.
func TestValidateStationDataParameter(t *testing.T) {
	t.Parallel()

	if !ValidateStationDataParameter("name") {
		t.Error("Unable to match field to slice member")
	}

	if !ValidateStationDataParameter("lum") {
		t.Error("Unable to match field to slice member")
	}

	if ValidateStationDataParameter("asdf") {
		t.Error("Found incorrect member in slice")
	}

	if !ValidateStationDataParameter("rain_1h") {
		t.Error("Found incorrect member in slice")
	}
}

// TestConvertToURLValues will make sure that ConvertToURLValues will
// convert a map[string]string to a url.Values instance and then to string.
func TestConvertToURLValues(t *testing.T) {
	t.Parallel()

	var count = 1
	var urlData = make(map[string]string)

	for _, s := range StationDataParameters {
		urlData[s] = string(count)
		count++
	}

	if reflect.TypeOf(ConvertToURLValues(urlData)).String() != "string" {
		t.Error("Unable to convert map to url.Values to string")
	}
}

// TestSendStationData will make sure that weather data will be sent to
// the OpenWeatherMap API.
func TestSendStationData(t *testing.T) {}
