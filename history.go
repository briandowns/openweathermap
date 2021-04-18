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
	OneH   float64 `json:"1h,omitempty"`
	ThreeH float64 `json:"3h"`
}

// Snow struct contains 3 hour data
type Snow struct {
	OneH   float64 `json:"1h,omitempty"`
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
}

// HistoryByName will return the history for the provided location.
func (o *OWM) HistoryByName(location string, hp *HistoricalParameters) (*HistoricalWeatherData, error) {
	base := fmt.Sprintf(historyURL, "appid=%s&%s&type=hour&mode=json&units=%s&lang=%s&cnt=%d&start=%d&end=%d")
	url := fmt.Sprintf(base, o.apiKey, "q="+url.QueryEscape(location), o.unit, o.lang, hp.Cnt, hp.Start, hp.End)

	var hwd HistoricalWeatherData
	if err := o.call(url, &hwd); err != nil {
		return nil, err
	}

	return &hwd, nil
}

// HistoryByID will return the history for the provided id.
func (o *OWM) HistoryByID(id int, hp *HistoricalParameters) (*HistoricalWeatherData, error) {
	base := fmt.Sprintf(historyURL, "appid=%s&id=%d&type=hour&mode=json&units=%s&lang=%s&cnt=%d&start=%d&end=%d")
	url := fmt.Sprintf(base, o.apiKey, id, o.unit, o.lang, hp.Cnt, hp.Start, hp.End)

	var hwd HistoricalWeatherData
	if err := o.call(url, &hwd); err != nil {
		return nil, err
	}

	return &hwd, nil
}

// HistoryByCoord will return the history for the provided coordinates
func (o *OWM) HistoryByCoord(location *Coordinates, hp *HistoricalParameters) (*HistoricalWeatherData, error) {
	base := fmt.Sprintf(historyURL, "appid=%s&lat=%f&lon=%f&start=%d&end=%d")
	url := fmt.Sprintf(base, o.apiKey, location.Latitude, location.Longitude, hp.Start, hp.End)

	var hwd HistoricalWeatherData
	if err := o.call(url, &hwd); err != nil {
		return nil, err
	}

	return &hwd, nil
}
