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

const (
	forecastFiveBase    = "https://api.openweathermap.org/data/2.5/forecast?appid=%s&%s&mode=json&units=%s&lang=%s&cnt=%d"
	forecastSixteenBase = "https://api.openweathermap.org/data/2.5/forecast/daily?appid=%s&%s&mode=json&units=%s&lang=%s&cnt=%d"
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
	Population int64       `json:"population"`
	Timezone   int64       `json:"timezome"`
	Sunrise    int64       `json:"sunrise"`
	Sunset     int64       `json:"sunset"`
}

// Forecast5Weather holds specific query data.
type ForecastFiveWeather struct {
	Dt         int64     `json:"dt"`
	Main       Main      `json:"main"`
	Weather    []Weather `json:"weather"`
	Clouds     Clouds    `json:"clouds"`
	Wind       Wind      `json:"wind"`
	Visibility int64     `json:"visibility"`
	Pop        float64   `json:"pop"`
	Rain       Rain      `json:"rain"`
	Sys        Sys       `json:"sys"`
	Snow       Snow      `json:"snow"`
	DtText     string    `json:"dt_text"`
}

// ForecastFiveWeatherData will hold returned data from queries.
type ForecastFiveWeatherData struct {
	COD     string                `json:"cod"`
	Message int64                 `json:"message"`
	Cnt     int64                 `json:"cnt"`
	List    []ForecastFiveWeather `json:"list"`
	City    City                  `json:"city"`
}

// ForecastSixteenWeather holds specific query data.
type ForecastSixteenWeather struct {
	Dt        int64            `json:"dt"`
	Temp      Temperature      `json:"temp"`
	Pressure  float64          `json:"pressure"`
	Humidity  int64            `json:"humidity"`
	Weather   []Weather        `json:"weather"`
	Speed     float64          `json:"speed"`
	Deg       int64            `json:"deg"`
	Clouds    int64            `json:"clouds"`
	Snow      float64          `json:"snow"`
	Rain      float64          `json:"rain"`
	FeelsLike FeelsLikeFullDay `json:"feels_like"`
	Pop       float64          `json:"pop"`
}

// ForecastSixteenWeatherData will hold returned data from queries.
type ForecastSixteenWeatherData struct {
	COD     string                   `json:"cod"`
	Message float64                  `json:"message"`
	City    City                     `json:"city"`
	Cnt     int64                    `json:"cnt"`
	List    []ForecastSixteenWeather `json:"list"`
}

// FiveDayForecastByName will provide a five day forecast for the given location.
func (o *OWM) FiveDayForecastByName(location string, cnt int) (*ForecastFiveWeatherData, error) {
	url := fmt.Sprintf(forecastFiveBase, o.apiKey, "q="+url.QueryEscape(location), o.unit, o.lang, cnt)

	var fwd ForecastFiveWeatherData
	if err := o.call(url, &fwd); err != nil {
		return nil, err
	}
	fmt.Printf("%#v\n", fwd)
	return &fwd, nil
}

// FiveDayForecastByCoordinates will provide a five day forecast for the given coordinates.
func (o *OWM) FiveDayForecastByCoordinates(location *Coordinates, cnt int) (*ForecastFiveWeatherData, error) {
	pos := fmt.Sprintf("lat=%f&lon=%f", location.Latitude, location.Longitude)
	url := fmt.Sprintf(forecastFiveBase, o.apiKey, pos, o.unit, o.lang, cnt)

	var fwd ForecastFiveWeatherData
	if err := o.call(url, &fwd); err != nil {
		return nil, err
	}

	return &fwd, nil
}

// FiveDayForecastByID will provide a forecast for the given location ID.
func (o *OWM) FiveDayForecastByID(id, cnt int) (*ForecastFiveWeatherData, error) {
	idq := fmt.Sprintf("id=%s", strconv.Itoa(id))
	url := fmt.Sprintf(forecastFiveBase, o.apiKey, idq, o.unit, o.lang, cnt)

	var fwd ForecastFiveWeatherData
	if err := o.call(url, &fwd); err != nil {
		return nil, err
	}

	return &fwd, nil
}

// FiveDayForecastByZip will provide a forecast for the given zip code.
func (o *OWM) FiveDayForecastByZip(zip, countryCode string, cnt int) (*ForecastFiveWeatherData, error) {
	zipq := fmt.Sprintf("zip=%s,%s", zip, countryCode)
	url := fmt.Sprintf(forecastFiveBase, o.apiKey, zipq, o.unit, o.lang, cnt)

	var fwd ForecastFiveWeatherData
	if err := o.call(url, &fwd); err != nil {
		return nil, err
	}

	return &fwd, nil
}

// SixteenDayForecastByName will provide a sixteen day forecast for the given location.
func (o *OWM) SixteenDayForecastByName(location string, cnt int) (*ForecastSixteenWeatherData, error) {
	url := fmt.Sprintf(forecastSixteenBase, o.apiKey, "q="+url.QueryEscape(location), o.unit, o.lang, cnt)

	var fwd ForecastSixteenWeatherData
	if err := o.call(url, &fwd); err != nil {
		return nil, err
	}
	fmt.Printf("%#v\n", fwd)
	return &fwd, nil
}

// SixteenDayForecastByCoordinates will provide a sixteen day forecast for the given coordinates.
func (o *OWM) SixteenDayForecastByCoordinates(location *Coordinates, cnt int) (*ForecastSixteenWeatherData, error) {
	pos := fmt.Sprintf("lat=%f&lon=%f", location.Latitude, location.Longitude)
	url := fmt.Sprintf(forecastSixteenBase, o.apiKey, pos, o.unit, o.lang, cnt)

	var fwd ForecastSixteenWeatherData
	if err := o.call(url, &fwd); err != nil {
		return nil, err
	}

	return &fwd, nil
}

// SixteenDayForecastByID will provide a forecast for the given location ID.
func (o *OWM) SixteenDayForecastByID(id, cnt int) (*ForecastSixteenWeatherData, error) {
	idq := fmt.Sprintf("id=%s", strconv.Itoa(id))
	url := fmt.Sprintf(forecastSixteenBase, o.apiKey, idq, o.unit, o.lang, cnt)

	var fwd ForecastSixteenWeatherData
	if err := o.call(url, &fwd); err != nil {
		return nil, err
	}

	return &fwd, nil
}

// SixteenDayForecastByZip will provide a forecast for the given zip code.
func (o *OWM) SixteenDayForecastByZip(zip, countryCode string, cnt int) (*ForecastSixteenWeatherData, error) {
	zipq := fmt.Sprintf("zip=%s,%s", zip, countryCode)
	url := fmt.Sprintf(forecastSixteenBase, o.apiKey, zipq, o.unit, o.lang, cnt)

	var fwd ForecastSixteenWeatherData
	if err := o.call(url, &fwd); err != nil {
		return nil, err
	}

	return &fwd, nil
}
