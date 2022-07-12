// Copyright 2022 Brian J. Downs
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
	"net/http"
	"strings"
)

var (
	errUnitUnavailable     = errors.New("unit unavailable")
	errLangUnavailable     = errors.New("language unavailable")
	errInvalidKey          = errors.New("invalid api key")
	errInvalidOption       = errors.New("invalid option")
	errInvalidHttpClient   = errors.New("invalid http client")
	errForecastUnavailable = errors.New("forecast unavailable")
	errExcludesUnavailable = errors.New("onecall excludes unavailable")
	errCountOfCityIDs      = errors.New("count of ids should not be more than 20 per request")
)

// DataUnits represents the character chosen to represent the temperature notation
var DataUnits = map[string]string{"C": "metric", "F": "imperial", "K": "internal"}
var (
	baseURL        = "https://api.openweathermap.org/data/2.5/weather?%s"
	onecallURL     = "https://api.openweathermap.org/data/2.5/onecall?%s"
	iconURL        = "https://openweathermap.org/img/w/%s"
	groupURL       = "http://api.openweathermap.org/data/2.5/group?%s"
	stationURL     = "https://api.openweathermap.org/data/2.5/station?id=%d"
	forecast5Base  = "https://api.openweathermap.org/data/2.5/forecast?appid=%s&%s&mode=json&units=%s&lang=%s&cnt=%d"
	forecast16Base = "https://api.openweathermap.org/data/2.5/forecast/daily?appid=%s&%s&mode=json&units=%s&lang=%s&cnt=%d"
	historyURL     = "https://api.openweathermap.org/data/2.5/history/%s"
	pollutionURL   = "https://api.openweathermap.org/data/2.5/air_pollution?appid=%s&lat=%s&lon=%s"
	uvURL          = "https://api.openweathermap.org/data/2.5/"
	dataPostURL    = "https://openweathermap.org/data/post"
)

// LangCodes holds all supported languages to be used
// inspried and sourced from @bambocher (github.com/bambocher)
var LangCodes = map[string]string{
	"AF":    "Afrikaans",
	"AL":    "Albanian",
	"AR":    "Arabic",
	"AZ":    "Azerbaijani",
	"BG":    "Bulgarian",
	"CA":    "Catalan",
	"CZ":    "Czech",
	"DA":    "Danish",
	"DE":    "German",
	"EL":    "Greek",
	"EN":    "English",
	"ES":    "Spanish",
	"EU":    "Basque",
	"FA":    "Persian (Farsi)",
	"FI":    "Finnish",
	"FR":    "French",
	"GL":    "Galician",
	"HE":    "Hebrew",
	"HI":    "Hindi",
	"HR":    "Croatian",
	"HU":    "Hungarian",
	"ID":    "Indonesian",
	"IT":    "Italian",
	"JA":    "Japanese",
	"KR":    "Korean",
	"LA":    "Latvian",
	"LT":    "Lithuanian",
	"MK":    "Macedonian",
	"NL":    "Dutch",
	"NO":    "Norwegian",
	"PL":    "Polish",
	"PT":    "Portuguese",
	"PT_BR": "Português Brasil",
	"RO":    "Romanian",
	"RU":    "Russian",
	"SE":    "Swedish",
	"SK":    "Slovak",
	"SL":    "Slovenian",
	"SP":    "Spanish",
	"SR":    "Serbian",
	"SV":    "Swedish",
	"TH":    "Thai",
	"TR":    "Turkish",
	"UA":    "Ukrainian",
	"UK":    "Ukrainian",
	"VI":    "Vietnamese",
	"ZH_CN": "Chinese Simplified",
	"ZH_TW": "Chinese Traditional",
	"ZU":    "Zulu",
}

// Exclude holds all supported excludes option to be used
const (
	ExcludeCurrent  = "current"
	ExcludeMinutely = "minutely"
	ExcludeHourly   = "hourly"
	ExcludeDaily    = "daily"
	ExcludeAlerts   = "alerts"
)

var Excludes []string = []string{
	ExcludeCurrent,
	ExcludeMinutely,
	ExcludeHourly,
	ExcludeDaily,
	ExcludeAlerts,
}

// Config will hold default settings to be passed into the
// "NewCurrent, NewForecast, etc}" functions.
type Config struct {
	Mode     string // user choice of JSON or XML
	Unit     string // measurement for results to be displayed.  F, C, or K
	Lang     string // should reference a key in the LangCodes map
	APIKey   string // API Key for connecting to the OWM
	Username string // Username for posting data
	Password string // Pasword for posting data
}

// APIError returned on failed API calls.
type APIError struct {
	Message string `json:"message"`
	COD     string `json:"cod"`
}

// Coordinates struct holds longitude and latitude data in returned
// JSON or as parameter data for requests using longitude and latitude.
type Coordinates struct {
	Longitude float64 `json:"lon"`
	Latitude  float64 `json:"lat"`
}

// Sys struct contains general information about the request
// and the surrounding area for where the request was made.
type Sys struct {
	Type    int     `json:"type"`
	ID      int     `json:"id"`
	Message float64 `json:"message"`
	Country string  `json:"country"`
	Sunrise int     `json:"sunrise"`
	Sunset  int     `json:"sunset"`
}

// Wind struct contains the speed and degree of the wind.
type Wind struct {
	Speed float64 `json:"speed"`
	Deg   float64 `json:"deg"`
}

// Weather struct holds high-level, basic info on the returned
// data.
type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

// Main struct contains the temperates, humidity, pressure for the request.
type Main struct {
	Temp      float64 `json:"temp"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	FeelsLike float64 `json:"feels_like"`
	Pressure  float64 `json:"pressure"`
	SeaLevel  float64 `json:"sea_level"`
	GrndLevel float64 `json:"grnd_level"`
	Humidity  int     `json:"humidity"`
}

// Clouds struct holds data regarding cloud cover.
type Clouds struct {
	All int `json:"all"`
}

// 	return key
// }
func setKey(key string) (string, error) {
	if err := ValidAPIKey(key); err != nil {
		return "", err
	}
	return key, nil
}

// ValidDataUnit makes sure the string passed in is an accepted
// unit of measure to be used for the return data.
func ValidDataUnit(u string) bool {
	for d := range DataUnits {
		if u == d {
			return true
		}
	}
	return false
}

// ValidLangCode makes sure the string passed in is an
// acceptable lang code.
func ValidLangCode(c string) bool {
	for d := range LangCodes {
		if c == d {
			return true
		}
	}
	return false
}

// ValidDataUnitSymbol makes sure the string passed in is an
// acceptable data unit symbol.
func ValidDataUnitSymbol(u string) bool {
	for _, d := range DataUnits {
		if u == d {
			return true
		}
	}
	return false
}

// ValidExcludes makes sure the string passed in is an
// acceptable excludes options.
func ValidExcludes(e []string) (string, error) {
	list := make([]string, 0)
	for _, v := range e {
		vl := strings.ToLower(v)
		notFound := true

		for _, d := range Excludes {
			if d == vl {
				list = append(list, v)
				notFound = false
				break
			}
		}

		if notFound {
			return "", errExcludesUnavailable
		}
	}
	return strings.Join(list, ","), nil
}

// ValidAPIKey makes sure that the key given is a valid one
func ValidAPIKey(key string) error {
	if len(key) > 64 {
		return errInvalidKey
	}
	return nil
}

// CheckAPIKeyExists will see if an API key has been set.
func (c *Config) CheckAPIKeyExists() bool { return len(c.APIKey) > 1 }

// Settings holds the client settings
type Settings struct {
	client *http.Client
}

// NewSettings returns a new Setting pointer with default http client.
func NewSettings() *Settings {
	return &Settings{
		client: http.DefaultClient,
	}
}

// Optional client settings
type Option func(s *Settings) error

// WithHttpClient sets custom http client when creating a new Client.
func WithHttpClient(c *http.Client) Option {
	return func(s *Settings) error {
		if c == nil {
			return errInvalidHttpClient
		}
		s.client = c
		return nil
	}
}

// setOptions sets Optional client settings to the Settings pointer
func setOptions(settings *Settings, options []Option) error {
	for _, option := range options {
		if option == nil {
			return errInvalidOption
		}
		err := option(settings)
		if err != nil {
			return err
		}
	}
	return nil
}
