// Copyright 2017 Gwennaël Buchet
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
	"net/url"
	"strconv"
	"strings"
)

// ForecastSys area population
type HourlyForecastCitySys struct {
	Population int `json:"population"`
}

type HourlyForecastMain struct {
	Temp      float64 `json:"temp"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  float64 `json:"pressure"`
	SeaLevel  float64 `json:"sea_level"`
	GrndLevel float64 `json:"grnd_level"`
	Humidity  float64 `json:"humidity"`
	TempKF    float64 `json:"temp_kf"`
}

type HourlyForecastWeather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type HourlyForecastClouds struct {
	All int `json:"all"`
}

type HourlyForecastWind struct {
	Speed float32 `json:"speed"`
	Deg   float32 `json:"deg"`
}

type HourlyForecastRain struct {
	ThreeH float32 `json:"3h"`
}

type HourlyForecastWeatherSys struct {
	Pod string `json:"pod"`
}

// HourlyForecastItem holds data for 1 time range
type HourlyForecastItem struct {
	Dt      int                      `json:"dt"`
	Main    HourlyForecastMain       `json:"main"`
	Weather []HourlyForecastWeather  `json:"weather"`
	Clouds  HourlyForecastClouds     `json:"clouds"`
	Wind    HourlyForecastWind       `json:"wind"`
	Rain    HourlyForecastRain       `json:"rain"`
	Sys     HourlyForecastWeatherSys `json:"sys"`
	DtTxt   string                   `json:"dt_txt"`
}

// City data for given location
type HourlyForecastCity struct {
	ID         int                   `json:"id"`
	Name       string                `json:"name"`
	Coord      Coordinates           `json:"coord"`
	Country    string                `json:"country"`
	Population int                   `json:"population"`
	Sys        HourlyForecastCitySys `json:"sys"`
}

// Global Hourly forecast data
type HourlyForecastData struct {
	City    HourlyForecastCity   `json:"city"`
	COD     string               `json:"cod"`
	Message float64              `json:"message"`
	CNT     int                  `json:"cnt"`
	List    []HourlyForecastItem `json:"list"`
	Unit    string
	Lang    string
	Key     string
	*Settings
}

// NewForecast returns a new HistoricalWeatherData pointer with
// the supplied arguments.
func NewHourlyForecast(unit, lang string, options ...Option) (*HourlyForecastData, error) {
	unitChoice := strings.ToUpper(unit)
	langChoice := strings.ToUpper(lang)

	f := &HourlyForecastData{
		Settings: NewSettings(),
	}

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

	if err := setOptions(f.Settings, options); err != nil {
		return nil, err
	}
	return f, nil
}

// HourlyForecastByName will provide a forecast for the location given for the
// number of days given.
func (f *HourlyForecastData) HourlyForecastByName(location string, days int) error {
	response, err := f.client.Get(fmt.Sprintf(forecastURL, f.Key, fmt.Sprintf("%s=%s", "q", url.QueryEscape(location)), f.Unit, f.Lang, days))

	if err != nil {
		return err
	}
	defer response.Body.Close()

	if err = json.NewDecoder(response.Body).Decode(&f); err != nil {
		return err
	}

	return nil
}

// HourlyForecastByCoordinates will provide a forecast for the coordinates ID give
// for the number of days given.
func (f *HourlyForecastData) HourlyForecastByCoordinates(location *Coordinates, days int) error {
	response, err := f.client.Get(fmt.Sprintf(forecastURL, f.Key, fmt.Sprintf("lat=%f&lon=%f", location.Latitude, location.Longitude), f.Unit, f.Lang, days))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if err = json.NewDecoder(response.Body).Decode(&f); err != nil {
		return err
	}

	return nil
}

// HourlyForecastByID will provide a forecast for the location ID give for the
// number of days given.
func (f *HourlyForecastData) HourlyForecastByID(id, days int) error {
	response, err := f.client.Get(fmt.Sprintf(forecastURL, f.Key, fmt.Sprintf("%s=%s", "id", strconv.Itoa(id)), f.Unit, f.Lang, days))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if err = json.NewDecoder(response.Body).Decode(&f); err != nil {
		return err
	}

	return nil
}
