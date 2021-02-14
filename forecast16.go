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
	"io"
)

// Forecast16WeatherList holds specific query data
type Forecast16WeatherList struct {
	Dt       int         `json:"dt"`
	Temp     Temperature `json:"temp"`
	Pressure float64     `json:"pressure"`
	Humidity int64       `json:"humidity"`
	Weather  []Weather   `json:"weather"`
	Speed    float64     `json:"speed"`
	Deg      int         `json:"deg"`
	Clouds   int         `json:"clouds"`
	Snow     float64     `json:"snow"`
	Rain     float64     `json:"rain"`
}

// Forecast16WeatherData will hold returned data from queries
type Forecast16WeatherData struct {
	COD     int                     `json:"cod"`
	Message string                  `json:"message"`
	City    City                    `json:"city"`
	Cnt     int                     `json:"cnt"`
	List    []Forecast16WeatherList `json:"list"`
}

func (f *Forecast16WeatherData) Decode(r io.Reader) error {
	if err := json.NewDecoder(r).Decode(&f); err != nil {
		return err
	}
	return nil
}
