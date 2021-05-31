// Copyright 2021 Marc-André Levasseur
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
	"strings"
)

var errExcludesInvalid = errors.New("invalid excludes options")

var (
	baseOnmecallURL = "https://api.openweathermap.org/data/2.5/onecall?%s"
)

type ExcludeOption int

const (
	ExcludeInvalid ExcludeOption = iota
	ExcludeCurrent
	ExcludeMinutely
	ExcludeHourly
	ExcludeDaily
	ExcludeAlerts
)

func (eop ExcludeOption) String() string {
	return map[ExcludeOption]string{
		ExcludeCurrent:  "current",
		ExcludeMinutely: "minutely",
		ExcludeHourly:   "hourly",
		ExcludeDaily:    "daily",
		ExcludeAlerts:   "alerts",
	}[eop]
}

type OneCallData struct {
	Latitude   float64          `json:"lat"`
	Longitude  float64          `json:"lon"`
	Timezone   string           `json:"timezone"`
	TimeOffset int              `json:"timezone_offset"`
	Current    OCCurrentData    `json:"current,omitempty"`
	Minutely   []OCMinutelyData `json:"minutely,omitempty"`
	Hourly     []OCHourlyData   `json:"hourly,omitempty"`
	Daily      []OCDailyData    `json:"daily,omitempty"`
	Alerts     []OCAlertData    `json:"alerts,omitempty"`
	// Internal settings
	Unit     string
	Lang     string
	Key      string
	Excludes string
	*Settings
}

type OCCurrentData struct {
	Dt          int       `json:"dt"`
	Sunrise     int       `json:"sunrise"`
	Sunset      int       `json:"sunset"`
	Temperature float64   `json:"feels_like"`
	FeelLike    float64   `json:"temp"`
	Pressure    float64   `json:"pressure"`
	Humidity    int       `json:"humidity"`
	DewPoint    float64   `json:"dew_point"`
	UVI         float64   `json:"uvi"`
	Clouds      int       `json:"clouds"`
	Visibility  int       `json:"visibility"`
	WindSpeed   float64   `json:"wind_speed"`
	WindDeg     float64   `json:"wind_deg"`
	WindGust    float64   `json:"wind_gust"`
	Weather     []Weather `json:"weather"`
}

type OCMinutelyData struct {
	Dt            int     `json:"dt"`
	Precipitation float64 `json:"precipitation"`
}

type OCHourlyData struct {
	Dt          int       `json:"dt"`
	Temperature float64   `json:"temp"`
	FeelsLike   float64   `json:"feels_like"`
	Pressure    float64   `json:"pressure"`
	Humidity    int       `json:"humidity"`
	DewPoint    float64   `json:"dew_point"`
	Uvi         float64   `json:"uvi"`
	Clouds      int       `json:"clouds"`
	Visibility  int       `json:"visibility"`
	WindSpeed   float64   `json:"wind_speed"`
	WindDeg     float64   `json:"wind_deg"`
	WindGust    float64   `json:"wind_gust"`
	Weather     []Weather `json:"weather"`
	Pop         float64   `json:"pop"`
}

type OCDailyData struct {
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
	Pressure  float64   `json:"pressure"`
	Humidity  int       `json:"humidity"`
	DewPoint  float64   `json:"dew_point"`
	WindSpeed float64   `json:"wind_speed"`
	WindDeg   float64   `json:"wind_deg"`
	WindGust  float64   `json:"wind_gust"`
	Weather   []Weather `json:"weather"`
	Clouds    int       `json:"clouds"`
	Pop       float64   `json:"pop"`
	Uvi       float64   `json:"uvi"`
}

type OCAlertData struct {
	SenderName  string `json:"sender_name"`
	Event       string `json:"event"`
	Start       int    `json:"start"`
	End         int    `json:"end"`
	Description string `json:"description"`
}

func NewOneCall(unit, lang, key string, exclude []ExcludeOption, options ...Option) (*OneCallData, error) {
	unitChoice := strings.ToUpper(unit)
	langChoice := strings.ToUpper(lang)

	oc := &OneCallData{
		Settings: NewSettings(),
	}

	if ValidDataUnit(unitChoice) {
		oc.Unit = DataUnits[unitChoice]
	} else {
		return nil, errUnitUnavailable
	}

	if ValidLangCode(langChoice) {
		oc.Lang = langChoice
	} else {
		return nil, errLangUnavailable
	}
	var err error
	ex, err := ValidateExcludes(exclude)
	if err != nil {
		return nil, err
	}
	oc.Excludes = ex

	oc.Key, err = setKey(key)
	if err != nil {
		return nil, err
	}

	if err := setOptions(oc.Settings, options); err != nil {
		return nil, err
	}
	return oc, nil
}

func (oc *OneCallData) PerformOneCall(latitude float64, longitude float64) error {
	params := fmt.Sprintf("appid=%s&lat=%f&lon=%f&units=%s&lang=%s", oc.Key, latitude, longitude, oc.Unit, oc.Lang)
	if oc.Excludes != "" {
		params = fmt.Sprintf("%s&exclude=%s", params, oc.Excludes)
	}
	response, err := oc.client.Get(fmt.Sprintf(baseOnmecallURL, params))
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if err = json.NewDecoder(response.Body).Decode(&oc); err != nil {
		return err
	}
	return nil
}

func ValidateExcludes(excludes []ExcludeOption) (string, error) {
	list := make([]string, 0)
	for _, e := range excludes {
		if e.String() == "" {
			return "", errExcludesInvalid
		}
		list = append(list, e.String())
	}
	return strings.Join(list, ","), nil
}
