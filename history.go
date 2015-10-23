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
	"strings"
)

// HistoricalParameters struct holds the (optional) fields to be
// supplied for historical data requests.
type HistoricalParameters struct {
	Start int64 // Data start (unix time, UTC time zone)
	End   int64 // Data end (unix time, UTC time zone)
	Cnt   int   // Amount of returned data (one per hour, can be used instead of Data end)
}

// Rain struct contains 3 hour data
type Rain struct {
	ThreeH float64 `json:"3h"`
}

// WeatherHistory struct contains aggregate fields from the above
// structs.
type WeatherHistory struct {
	Main    Main      `json:"main"`
	Wind    Wind      `json:"wind"`
	Clouds  Clouds    `json:"clouds"`
	Weather []Weather `json:"weather"`
	Rain    Rain      `json:"rain"`
	Dt      int       `json:"dt"`
}

// HistoricalWeatherData struct is where the JSON is unmarshaled to
// when receiving data for a historical request.
type HistoricalWeatherData struct {
	Message  string           `json:"message"`
	Cod      int              `json:"cod"`
	CityData int              `json:"city_data"`
	CalcTime float64          `json:"calctime"`
	Cnt      int              `json:"cnt"`
	List     []WeatherHistory `json:"list"`
	Unit     string
}

// NewHistorical returns a new HistoricalWeatherData pointer with
//the supplied arguments.
func NewHistorical(unit string) (*HistoricalWeatherData, error) {
	unitChoice := strings.ToUpper(unit)
	if ValidDataUnit(unitChoice) {
		return &HistoricalWeatherData{Unit: DataUnits[unitChoice]}, nil
	}
	return nil, errors.New("unit of measure not available")
}

// HistoryByName will return the history for the provided location
func (h *HistoricalWeatherData) HistoryByName(location string) error {
	var err error
	var response *http.Response
	if !config.CheckAPIKeyExists() {
		return ErrApiKeyNotFound
	}
	response, err = http.Get(fmt.Sprintf(fmt.Sprintf(historyURL, "city?q=%s&appid=%s"), url.QueryEscape(location), config.APIKey))
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if err = json.NewDecoder(response.Body).Decode(&h); err != nil {
		return err
	}
	return nil
}

// HistoryByID will return the history for the provided location ID
func (h *HistoricalWeatherData) HistoryByID(id int, hp ...*HistoricalParameters) error {
	if !config.CheckAPIKeyExists() {
		return ErrApiKeyNotFound
	}
	if len(hp) > 0 {
		response, err := http.Get(fmt.Sprintf(fmt.Sprintf(historyURL, "city?id=%d&type=hour&start%d&end=%d&cnt=%d&appid=%s"), id, hp[0].Start, hp[0].End, hp[0].Cnt, config.APIKey))
		if err != nil {
			return err
		}
		defer response.Body.Close()
		if err = json.NewDecoder(response.Body).Decode(&h); err != nil {
			return err
		}
	}
	response, err := http.Get(fmt.Sprintf(fmt.Sprintf(historyURL, "city?id=%d&appid=%s"), id, config.APIKey))
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if err = json.NewDecoder(response.Body).Decode(&h); err != nil {
		return err
	}
	return nil
}
