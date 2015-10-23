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
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Forecast5D3HList holds specific query data
type Forecast3HList struct {
	Dt      int       `json:"dt"`
	Main    Main      `json:"main"`
	Weather []Weather `json:"weather"`
	Clouds  Clouds    `json:"clouds"`
	Wind    Wind      `json:"wind"`
	Rain    Rain      `json:"rain"`
	Snow    Snow      `json:"snow"`
	DtTxt   string    `json:"dt_txt "`
}

// ForecastWeatherData will hold returned data from queries
type Forecast3HData struct {
	COD     string           `json:"cod"`
	Message float64          `json:"message"`
	City    City             `json:"city"`
	Cnt     int              `json:"cnt"`
	List    []Forecast3HList `json:"list"`
	Unit    string
	Lang    string
}

// NewForecast returns a new HistoricalWeatherData pointer with
// the supplied arguments.
func NewForecast3H(unit, lang string) (*Forecast3HData, error) {
	unitChoice := strings.ToUpper(unit)
	langChoice := strings.ToUpper(lang)
	f := &Forecast3HData{}
	if ValidDataUnit(unitChoice) {
		f.Unit = DataUnits[unitChoice]
	} else {
		return nil, errors.New(unitError)
	}
	if ValidLangCode(langChoice) {
		f.Lang = langChoice
	} else {
		return nil, errors.New(langError)
	}
	return f, nil
}

// DailyByName will provide a forecast for the location given for the
// number of days given.
func (f *Forecast3HData) ByName(location string, days int) error {
	var err error
	var response *http.Response
	if !config.CheckAPIKeyExists() {
		return ErrApiKeyNotFound
	}
	response, err = http.Get(fmt.Sprintf(forecast3HBase, config.APIKey, "q", url.QueryEscape(location), f.Unit, f.Lang, days))
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if err = json.NewDecoder(response.Body).Decode(&f); err != nil {
		return err
	}
	return nil
}

// ByCoordinates will provide a forecast for the coordinates ID give
// for the number of days given.
func (f *Forecast3HData) ByCoordinates(location *Coordinates, days int) error {
	if !config.CheckAPIKeyExists() {
		return ErrApiKeyNotFound
	}
	response, err := http.Get(fmt.Sprintf(fmt.Sprintf(forecast3HBase, config.APIKey, "lat=%f&lon=%f&units=%s"), location.Latitude, location.Longitude, f.Unit, f.Lang, days))
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if err = json.NewDecoder(response.Body).Decode(&f); err != nil {
		return err
	}
	return nil
}

// ByID will provide a forecast for the location ID give for the
// number of days given.
func (f *Forecast3HData) ByID(id, days int) error {
	if !config.CheckAPIKeyExists() {
		return ErrApiKeyNotFound
	}
	response, err := http.Get(fmt.Sprintf(forecast3HBase, config.APIKey, "id", strconv.Itoa(id), f.Unit, f.Lang, days))
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if err = json.NewDecoder(response.Body).Decode(&f); err != nil {
		return err
	}
	return nil
}
