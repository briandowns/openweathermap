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
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
)

// CurrentWeatherData struct contains an aggregate view of the structs
// defined above for JSON to be unmarshaled into.
type CurrentWeatherData struct {
	GeoPos   Coordinates `json:"coord"`
	Sys      Sys         `json:"sys"`
	Base     string      `json:"base"`
	Weather  []Weather   `json:"weather"`
	Main     Main        `json:"main"`
	Wind     Wind        `json:"wind"`
	Clouds   Clouds      `json:"clouds"`
	Rain     Rain        `json:"rain"`
	Snow     Snow        `json:"snow"`
	Dt       int         `json:"dt"`
	ID       int         `json:"id"`
	Name     string      `json:"name"`
	Cod      int         `json:"cod"`
	Timezone int         `json:"timezone"`
	Unit     string
	Lang     string
	Key      string
}

// CurrentByName will provide the current weather with
// the provided location name.
func (o *OWM) CurrentByName(location string) (*CurrentWeatherData, error) {
	base := fmt.Sprintf(baseURL, "appid=%s&q=%s&units=%s&lang=%s")
	url := fmt.Sprintf(base, o.apiKey, url.QueryEscape(location), o.unit, o.lang)

	var cwd CurrentWeatherData
	if err := o.call(url, &cwd); err != nil {
		return nil, err
	}

	return &cwd, nil
}

// CurrentByCoordinates will provide the current weather
// with the provided location coordinates.
func (o *OWM) CurrentByCoordinates(location *Coordinates) (*CurrentWeatherData, error) {
	base := fmt.Sprintf(baseURL, "appid=%s&lat=%f&lon=%f&units=%s&lang=%s")
	url := fmt.Sprintf(base, o.apiKey, location.Latitude, location.Longitude, o.unit, o.lang)

	var cwd CurrentWeatherData
	if err := o.call(url, &cwd); err != nil {
		return nil, err
	}

	return &cwd, nil
}

// CurrentByID will provide the current weather with the
// provided location ID.
func (o *OWM) CurrentByID(id int) (*CurrentWeatherData, error) {
	base := fmt.Sprintf(baseURL, "appid=%s&id=%d&units=%s&lang=%s")
	url := fmt.Sprintf(base, o.apiKey, id, o.unit, o.lang)

	var cwd CurrentWeatherData
	if err := o.call(url, &cwd); err != nil {
		return nil, err
	}

	return &cwd, nil
}

// CurrentByZip will provide the current weather for the
// provided zip code.
func (o *OWM) CurrentByZip(zip int, countryCode string) (*CurrentWeatherData, error) {
	base := fmt.Sprintf(baseURL, "appid=%s&zip=%d,%s&units=%s&lang=%s")
	url := fmt.Sprintf(base, o.apiKey, zip, countryCode, o.unit, o.lang)

	var cwd CurrentWeatherData
	if err := o.call(url, &cwd); err != nil {
		return nil, err
	}

	return &cwd, nil
}

// CurrentByArea will provide the current weather for the
// provided area.
func (o *OWM) CurrentByArea() (*CurrentWeatherData, error) {
	return nil, errors.New("unimplemented")
}

func (o *OWM) call(url string, payload interface{}) error {
	res, err := o.client.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(payload)
}
