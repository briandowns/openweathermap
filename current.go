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
	"io/ioutil"
	"net/http"
	"strings"
)

// CurrentWeatherData struct contains an aggregate view of the structs
// defined above for JSON to be unmarshaled into.
type CurrentWeatherData struct {
	GeoPos  Coordinates `json:"coord"`
	Sys     Sys         `json:"sys"`
	Base    string      `json:"base"`
	Weather []Weather   `json:"weather"`
	Main    Main        `json:"main"`
	Wind    Wind        `json:"wind"`
	Clouds  Clouds      `json:"clouds"`
	Dt      int         `json:"dt"`
	ID      int         `json:"id"`
	Name    string      `json:"name"`
	Cod     int         `json:"cod"`
	Units   string
}

// NewCurrent returns a new WeatherData pointer with the supplied.
func NewCurrent(unit string) (*CurrentWeatherData, error) {
	unitChoice := strings.ToLower(unit)
	if ValidDataUnit(unitChoice) {
		return &CurrentWeatherData{Units: unitChoice}, nil
	}
	return nil, errors.New("ERROR: unit of measure not available")
}

// CurrentByName will provide the current weather with the
// provided location name.
func (w *CurrentWeatherData) CurrentByName(location string) error {
	response, err := http.Get(fmt.Sprintf(fmt.Sprintf(baseURL, "q=%s&units=%s"), location, w.Units))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(result, &w)
	if err != nil {
		return err
	}
	return nil
}

// CurrentByCoordinates will provide the current weather with the
// provided location coordinates.
func (w *CurrentWeatherData) CurrentByCoordinates(location *Coordinates) error {
	response, err := http.Get(fmt.Sprintf(fmt.Sprintf(baseURL, "lat=%f&lon=%f&units=%s"), location.Latitude, location.Longitude, w.Units))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(result, &w)
	if err != nil {
		return err
	}
	return nil
}

// CurrentByID will provide the current weather with the
// provided location ID.
func (w *CurrentWeatherData) CurrentByID(id int) error {
	response, err := http.Get(fmt.Sprintf(fmt.Sprintf(baseURL, "id=%d&units=%s"), id, w.Units))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(result, &w)
	if err != nil {
		return err
	}
	return nil
}

// CurrentByArea will provide the current weather for the
// provided area.
func (w *CurrentWeatherData) CurrentByArea() {}
