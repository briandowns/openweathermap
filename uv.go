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
	"errors"
	"fmt"
	"time"
)

// UVDataPoints holds the UV specific data.
type UVDataPoints struct {
	DT    int64   `json:"dt"`
	Value float64 `json:"value"`
}

// UV contains the response from the OWM UV API.
type UV struct {
	Coordinates
	DateISO string  `json:"date_iso"`
	Date    int64   `json:"date,omitempty"`
	Value   float64 `json:"value,omitempty"`
}

// UVCurrent gets the current UV data for the given coordinates.
func (o *OWM) UVCurrent(coord *Coordinates) (*UV, error) {
	url := fmt.Sprintf("%suvi?lat=%f&lon=%f&appid=%s", uvURL, coord.Latitude, coord.Longitude, o.apiKey)

	var uv UV
	if err := o.call(url, &uv); err != nil {
		return nil, err
	}

	return &uv, nil
}

// UVHistorical gets the historical UV data for the coordinates and times.
func (o *OWM) UVHistorical(coord *Coordinates, start, end time.Time) (*UV, error) {
	url := fmt.Sprintf("%shistory?lat=%f&lon=%f&start=%d&end=%d&appid=%s", uvURL, coord.Latitude, coord.Longitude, start.Unix(), end.Unix(), o.apiKey)

	var uv UV
	if err := o.call(url, &uv); err != nil {
		return nil, err
	}

	return &uv, nil
}

// UVIndexInfo
type UVIndexInfo struct {
	// UVIndex holds the range of the index
	UVIndex []float64

	// MGC represents the Media graphic color
	MGC string

	// Risk of harm from unprotected sun exposure, for the average adult
	Risk string

	// RecommendedProtection contains information on what a person should
	// do when outside in the associated UVIndex
	RecommendedProtection string
}

// UVData contains data in regards to UV index ranges, rankings, and steps for protection.
var UVData = []UVIndexInfo{
	{
		UVIndex:               []float64{0, 2.9},
		MGC:                   "Green",
		Risk:                  "Low",
		RecommendedProtection: "Wear sunglasses on bright days; use sunscreen if there is snow on the ground, which reflects UV radiation, or if you have particularly fair skin.",
	},
	{
		UVIndex:               []float64{3, 5.9},
		MGC:                   "Yellow",
		Risk:                  "Moderate",
		RecommendedProtection: "Take precautions, such as covering up, if you will be outside. Stay in shade near midday when the sun is strongest.",
	},
	{
		UVIndex:               []float64{6, 7.9},
		MGC:                   "Orange",
		Risk:                  "High",
		RecommendedProtection: "Cover the body with sun protective clothing, use SPF 30+ sunscreen, wear a hat, reduce time in the sun within three hours of solar noon, and wear sunglasses.",
	},
	{
		UVIndex:               []float64{8, 10.9},
		MGC:                   "Red",
		Risk:                  "Very high",
		RecommendedProtection: "Wear SPF 30+ sunscreen, a shirt, sunglasses, and a wide-brimmed hat. Do not stay in the sun for too long.",
	},
	{
		UVIndex:               []float64{11},
		MGC:                   "Violet",
		Risk:                  "Extreme",
		RecommendedProtection: "Take all precautions: Wear SPF 30+ sunscreen, a long-sleeved shirt and trousers, sunglasses, and a very broad hat. Avoid the sun within three hours of solar noon.",
	},
}

// UVInformation provides information on the given UV data which includes the severity
// and "Recommended protection"
func (u *UV) UVInformation() ([]UVIndexInfo, error) {
	switch {
	case u.Value != 0:
		switch {
		case u.Value < 2.9:
			return []UVIndexInfo{UVData[0]}, nil
		case u.Value > 3 && u.Value < 5.9:
			return []UVIndexInfo{UVData[1]}, nil
		case u.Value > 6 && u.Value < 7.9:
			return []UVIndexInfo{UVData[2]}, nil
		case u.Value > 8 && u.Value < 10.9:
			return []UVIndexInfo{UVData[3]}, nil
		case u.Value >= 11:
			return []UVIndexInfo{UVData[4]}, nil
		default:
			return nil, errors.New("invalid UV index value")
		}
	}

	return nil, nil
}
