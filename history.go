// Copyright 2014 Brian J. Downs
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
	"strings"
)

// HistoricalParameters struct holds the (optional) fields to be
// supplied for historical data requests.
type HistoricalParameters struct {
	Start int64
	End   int64
	Cnt   int
}

type Rain struct {
	threeH int `json:"3h"`
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
		return &HistoricalWeatherData{Unit: unitChoice}, nil
	}
	return nil, errors.New("unit of measure not available")
}

func (h *HistoricalWeatherData) HistoryByName(location string) error {
	response, err := http.Get(fmt.Sprintf(fmt.Sprintf(historyURL, "city?q=%s"), location))
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if err = json.NewDecoder(response.Body).Decode(&h); err != nil {
		return err
	}
	return nil
}

func (h *HistoricalWeatherData) HistoryByCoordinates(location *Coordinates) error {
	return nil
}

func (h *HistoricalWeatherData) HistoryByID(id int) error {
	return nil
}

func (h *HistoricalWeatherData) HistoryByArea() error {
	return nil
}
