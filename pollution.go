// Copyright 2021 Brian J. Downs
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
	"fmt"
	"strconv"
)

// DateTimeAliases holds the alias the pollution API supports in lieu
// of an ISO 8601 timestamp
var DateTimeAliases = []string{"current"}

// ValidAlias checks to make sure the given alias is a valid one
func ValidAlias(alias string) bool {
	for _, i := range DateTimeAliases {
		if i == alias {
			return true
		}
	}
	return false
}

// PollutionData holds the pollution specific data from the call
type PollutionData struct {
	Precision float64 `json:"precision"`
	Pressure  float64 `json:"pressure"`
	Value     float64 `json:"value"`
}

// PollutionParameters holds the parameters needed to make
// a call to the pollution API
type PollutionParameters struct {
	Location Coordinates
	Datetime string // this should be either ISO 8601 or an alias
}

// Pollution holds the data returnd from the pollution API
type Pollution struct {
	Time     string          `json:"time"`
	Location Coordinates     `json:"location"`
	Data     []PollutionData `json:"data"`
	Key      string
}

// PollutionByParams gets the pollution data based on the given parameters
func (o *OWM) PollutionByParams(params *PollutionParameters) (*Pollution, error) {
	url := fmt.Sprintf("%s%s,%s/%s.json?appid=%s",
		pollutionURL,
		strconv.FormatFloat(params.Location.Latitude, 'f', -1, 64),
		strconv.FormatFloat(params.Location.Longitude, 'f', -1, 64),
		params.Datetime,
		o.apiKey)

	var p Pollution
	if err := o.call(url, &p); err != nil {
		return nil, err
	}

	return &p, nil
}
