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
	"strings"
	"time"
)

type DtTxt struct {
	time.Time
}

func (dt *DtTxt) UnmarshalJSON(b []byte) error {
	t, err := time.Parse("2006-01-02 15:04:05", strings.Trim(string(b), "\""))
	dt.Time = t
	return err
}

func (t *DtTxt) MarshalJSON() ([]byte, error) {
	return json.Marshal(t)
}

// Forecast5WeatherList holds specific query data
type Forecast5WeatherList struct {
	Dt      int       `json:"dt"`
	Main    Main      `json:"main"`
	Weather []Weather `json:"weather"`
	Clouds  Clouds    `json:"clouds"`
	Wind    Wind      `json:"wind"`
	Rain    Rain      `json:"rain"`
	Snow    Snow      `json:"snow"`
	DtTxt   DtTxt     `json:"dt_txt"`
}

// Forecast5WeatherData will hold returned data from queries
type Forecast5WeatherData struct {
	// COD     string                `json:"cod"`
	// Message float64               `json:"message"`
	City City                   `json:"city"`
	Cnt  int                    `json:"cnt"`
	List []Forecast5WeatherList `json:"list"`
}

func (f *Forecast5WeatherData) Decode(r io.Reader) error {
	if err := json.NewDecoder(r).Decode(&f); err != nil {
		return err
	}
	return nil
}
