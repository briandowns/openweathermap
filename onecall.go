// Copyright 2022 Giuseppe Silvestro
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
	"strings"
)

// OneCallData struct contains an aggregate view of the structs
// defined above for JSON to be unmarshaled into.
type OneCallData struct {
	Latitude       float64               `json:"lat"`
	Longitude      float64               `json:"lon"`
	Timezone       string                `json:"timezone"`
	TimezoneOffset int                   `json:"timezone_offset"`
	Current        OneCallCurrentData    `json:"current,omitempty"`
	Minutely       []OneCallMinutelyData `json:"minutely,omitempty"`
	Hourly         []OneCallHourlyData   `json:"hourly,omitempty"`
	Daily          []OneCallDailyData    `json:"daily,omitempty"`
	Alerts         []OneCallAlertData    `json:"alerts,omitempty"`

	Unit     string
	Lang     string
	Key      string
	Excludes string
	*Settings
}

type OneCallCurrentData struct {
	Dt         int       `json:"dt"`
	Sunrise    int       `json:"sunrise"`
	Sunset     int       `json:"sunset"`
	Temp       float64   `json:"temp"`
	FeelsLike  float64   `json:"feels_like"`
	Pressure   int       `json:"pressure"`
	Humidity   int       `json:"humidity"`
	DewPoint   float64   `json:"dew_point"`
	Clouds     int       `json:"clouds"`
	UVI        float64   `json:"uvi"`
	Visibility int       `json:"visibility"`
	WindSpeed  float64   `json:"wind_speed"`
	WindGust   float64   `json:"wind_gust,omitempty"`
	WindDeg    float64   `json:"wind_deg"`
	Rain       Rain      `json:"rain,omitempty"`
	Snow       Snow      `json:"snow,omitempty"`
	Weather    []Weather `json:"weather"`
}

type OneCallMinutelyData struct {
	Dt            int     `json:"dt"`
	Precipitation float64 `json:"precipitation"`
}

type OneCallHourlyData struct {
	Dt         int       `json:"dt"`
	Temp       float64   `json:"temp"`
	FeelsLike  float64   `json:"feels_like"`
	Pressure   int       `json:"pressure"`
	Humidity   int       `json:"humidity"`
	DewPoint   float64   `json:"dew_point"`
	UVI        float64   `json:"uvi"`
	Clouds     int       `json:"clouds"`
	Visibility int       `json:"visibility"`
	WindSpeed  float64   `json:"wind_speed"`
	WindGust   float64   `json:"wind_gust,omitempty"`
	WindDeg    float64   `json:"wind_deg"`
	Pop        float64   `json:"pop"`
	Rain       Rain      `json:"rain,omitempty"`
	Snow       Snow      `json:"snow,omitempty"`
	Weather    []Weather `json:"weather"`
}

type OneCallDailyData struct {
	Dt        int         `json:"dt"`
	Sunrise   int         `json:"sunrise"`
	Sunset    int         `json:"sunset"`
	Moonrise  int         `json:"moonrise"`
	Moonset   int         `json:"moonset"`
	MoonPhase float64     `json:"moon_phase"`
	Temp      Temperature `json:"temp"`
	FeelsLike struct {
		Day   float64 `json:"day"`
		Night float64 `json:"night"`
		Eve   float64 `json:"eve"`
		Morn  float64 `json:"morn"`
	} `json:"feels_like"`
	Pressure  int       `json:"pressure"`
	Humidity  int       `json:"humidity"`
	DewPoint  float64   `json:"dew_point"`
	WindSpeed float64   `json:"wind_speed"`
	WindGust  float64   `json:"wind_gust,omitempty"`
	WindDeg   float64   `json:"wind_deg"`
	Clouds    int       `json:"clouds"`
	UVI       float64   `json:"uvi"`
	Pop       float64   `json:"pop"`
	Rain      float64   `json:"rain,omitempty"`
	Snow      float64   `json:"snow,omitempty"`
	Weather   []Weather `json:"weather"`
}

type OneCallAlertData struct {
	SenderName  string   `json:"sender_name"`
	Event       string   `json:"event"`
	Start       int      `json:"start"`
	End         int      `json:"end"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

// NewCurrent returns a new OneCallData pointer with the supplied parameters
func NewOneCall(unit, lang, key string, excludes []string, options ...Option) (*OneCallData, error) {
	unitChoice := strings.ToUpper(unit)
	langChoice := strings.ToUpper(lang)

	c := &OneCallData{
		Settings: NewSettings(),
	}

	if !ValidDataUnit(unitChoice) {
		return nil, errUnitUnavailable
	}
	c.Unit = DataUnits[unitChoice]

	if !ValidLangCode(langChoice) {
		return nil, errLangUnavailable
	}
	c.Lang = langChoice

	var err error
	c.Excludes, err = ValidExcludes(excludes)
	if err != nil {
		return nil, err
	}

	c.Key, err = setKey(key)
	if err != nil {
		return nil, err
	}

	err = setOptions(c.Settings, options)
	if err != nil {
		return nil, err
	}

	return c, nil
}

// OneCallByCoordinates will provide the onecall weather with the
// provided location coordinates.
func (w *OneCallData) OneCallByCoordinates(location *Coordinates) error {
	response, err := w.client.Get(fmt.Sprintf(fmt.Sprintf(onecallURL, "appid=%s&lat=%f&lon=%f&units=%s&lang=%s&exclude=%s"), w.Key, location.Latitude, location.Longitude, w.Unit, w.Lang, w.Excludes))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	return json.NewDecoder(response.Body).Decode(&w)
}
