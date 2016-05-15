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
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// ForecastSys area population
type ForecastSys struct {
	Population int `json:"population"`
}

// Temperature holds returned termperate sure stats
type Temperature struct {
	Day   float64 `json:"day"`
	Min   float64 `json:"min"`
	Max   float64 `json:"max"`
	Night float64 `json:"night"`
	Eve   float64 `json:"eve"`
	Morn  float64 `json:"morn"`
}

// City data for given location
type City struct {
	ID         int         `json:"id"`
	Name       string      `json:"name"`
	Coord      Coordinates `json:"coord"`
	Country    string      `json:"country"`
	Population int         `json:"population"`
	Sys        ForecastSys `json:"sys"`
}

// ForecastWeatherList holds specific query data
type ForecastWeatherList struct {
	Dt       int         `json:"dt"`
	Temp     Temperature `json:"temp"`
	Pressure float64     `json:"pressure"`
	Humidity int         `json:"humidity"`
	Weather  []Weather   `json:"weather"`
	Speed    float64     `json:"speed"`
	Deg      int         `json:"deg"`
	Clouds   int         `json:"clouds"`
	Snow     float64     `json:"snow"`
	Rain     float64     `json:"rain"`
}

// ForecastWeatherData will hold returned data from queries
type ForecastWeatherData struct {
	COD     string                `json:"cod"`
	Message float64               `json:"message"`
	City    City                  `json:"city"`
	Cnt     int                   `json:"cnt"`
	List    []ForecastWeatherList `json:"list"`
	Unit    string
	Lang    string
	Key     string
}

// NewForecast returns a new HistoricalWeatherData pointer with
// the supplied arguments.
func NewForecast(unit, lang string) (*ForecastWeatherData, error) {
	unitChoice := strings.ToUpper(unit)
	langChoice := strings.ToUpper(lang)

	f := &ForecastWeatherData{}

	if ValidDataUnit(unitChoice) {
		f.Unit = DataUnits[unitChoice]
	} else {
		return nil, errUnitUnavailable
	}

	if ValidLangCode(langChoice) {
		f.Lang = langChoice
	} else {
		return nil, errLangUnavailable
	}

	f.Key = getKey()

	return f, nil
}

// DailyByName will provide a forecast for the location given for the
// number of days given.
func (f *ForecastWeatherData) DailyByName(location string, days int) error {
	var err error
	var response *http.Response

	response, err = http.Get(fmt.Sprintf(forecastBase, f.Key, fmt.Sprintf("%s=%s","q", url.QueryEscape(location)), f.Unit, f.Lang, days))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if err = json.NewDecoder(response.Body).Decode(&f); err != nil {
		return err
	}

	return nil
}

// DailyByCoordinates will provide a forecast for the coordinates ID give
// for the number of days given.
func (f *ForecastWeatherData) DailyByCoordinates(location *Coordinates, days int) error {
	response, err := http.Get(fmt.Sprintf(forecastBase, f.Key, fmt.Sprintf("lat=%f&lon=%f", location.Latitude, location.Longitude), f.Unit, f.Lang, days))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if err = json.NewDecoder(response.Body).Decode(&f); err != nil {
		return err
	}

	return nil
}

// DailyByID will provide a forecast for the location ID give for the
// number of days given.
func (f *ForecastWeatherData) DailyByID(id, days int) error {
	response, err := http.Get(fmt.Sprintf(forecastBase, f.Key, fmt.Sprintf("%s=%s","id", strconv.Itoa(id)), f.Unit, f.Lang, days))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if err = json.NewDecoder(response.Body).Decode(&f); err != nil {
		return err
	}

	return nil
}
