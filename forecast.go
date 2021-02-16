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
	forecastBase = "https://api.openweathermap.org/data/2.5/forecast?appid=%s&%s&mode=json&units=%s&lang=%s&cnt=%d"
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
	COD string `json:"cod"`
	//Message int64  `json:"message"`
	//Cnt     int64  `json:"cnt"`
	//List    []ForecastFiuveWeather `json:"list"`
	City City `json:"city"`
}

// FiveDayForecastByName will provide a five day forecast for the given location.
func (o *OWM) FiveDayForecastByName(location string, cnt int) (*ForecastFiveWeatherData, error) {
	url := fmt.Sprintf(forecastBase, o.apiKey, "q="+url.QueryEscape(location), o.unit, o.lang, cnt)
	fmt.Println(url)
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
	url := fmt.Sprintf(forecastBase, o.apiKey, pos, o.unit, o.lang, cnt)
	fmt.Println(url)
	var fwd ForecastFiveWeatherData
	if err := o.call(url, &fwd); err != nil {
		return nil, err
	}

	return &fwd, nil
}

// DailyForecastByID will provide a forecast for the location ID give for the
// number of days given.
func (o *OWM) DailyForecastByID(id, days int) (*ForecastFiveWeatherData, error) {
	base := fmt.Sprintf("%s=%s", "id", strconv.Itoa(id))
	url := fmt.Sprintf(base, o.apiKey, o.unit, o.lang, days)

	var fwd ForecastFiveWeatherData
	if err := o.call(url, &fwd); err != nil {
		return nil, err
	}

	return &fwd, nil
}

// DailyForecastByZip will provide a forecast for the provided zip code.
func (o *OWM) DailyForecastByZip(zip, countryCode string, days int) (*ForecastFiveWeatherData, error) {
	base := fmt.Sprintf("zip=%s,%s", zip, countryCode)
	url := fmt.Sprintf(base, o.apiKey, o.unit, o.lang, days)

	var fwd ForecastFiveWeatherData
	if err := o.call(url, &fwd); err != nil {
		return nil, err
	}

	return &fwd, nil
}
