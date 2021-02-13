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
	"net/url"
	"strconv"
)

// ForecastSys area population.
type ForecastSys struct {
	Population int `json:"population"`
}

// Temperature holds returned termperate sure stats.
type Temperature struct {
	Day   float64 `json:"day"`
	Min   float64 `json:"min"`
	Max   float64 `json:"max"`
	Night float64 `json:"night"`
	Eve   float64 `json:"eve"`
	Morn  float64 `json:"morn"`
}

// City data for given location.
type City struct {
	ID         int         `json:"id"`
	Name       string      `json:"name"`
	Coord      Coordinates `json:"coord"`
	Country    string      `json:"country"`
	Population int         `json:"population"`
	Sys        ForecastSys `json:"sys"`
}

// DailyForecastByName will provide a forecast for the location given for the
// number of days given.
func (o *OWM) DailyForecastByName(location string, days int) (interface{}, error) {
	base := fmt.Sprintf("%s=%s", "q", url.QueryEscape(location))
	url := fmt.Sprintf(base, o.apiKey, o.unit, o.lang, days)

	var fwd interface{}
	if err := o.call(url, &fwd); err != nil {
		return nil, err
	}

	return &fwd, nil
}

// DailyForecastByCoordinates will provide a forecast for the coordinates ID give
// for the number of days given.
func (o *OWM) DailyForecastByCoordinates(location *Coordinates, days int) (interface{}, error) {
	base := fmt.Sprintf("lat=%f&lon=%f", location.Latitude, location.Longitude)
	url := fmt.Sprintf(base, o.apiKey, o.unit, o.lang, days)

	var fwd interface{}
	if err := o.call(url, &fwd); err != nil {
		return nil, err
	}

	return &fwd, nil
}

// DailyForecastByID will provide a forecast for the location ID give for the
// number of days given.
func (o *OWM) DailyForecastByID(id, days int) (interface{}, error) {
	base := fmt.Sprintf("%s=%s", "id", strconv.Itoa(id))
	url := fmt.Sprintf(base, o.apiKey, o.unit, o.lang, days)

	var fwd interface{}
	if err := o.call(url, &fwd); err != nil {
		return nil, err
	}

	return &fwd, nil
}

// DailyForecastByZip will provide a forecast for the provided zip code.
func (o *OWM) DailyForecastByZip(zip, countryCode string, days int) (interface{}, error) {
	base := fmt.Sprintf("zip=%s,%s", zip, countryCode)
	url := fmt.Sprintf(base, o.apiKey, o.unit, o.lang, days)

	var fwd interface{}
	if err := o.call(url, &fwd); err != nil {
		return nil, err
	}

	return &fwd, nil
}
