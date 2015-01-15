package openweathermap

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// ForecastSys area population
type ForecastSys struct {
	Population int `json:"population"`
}

// Temperature holds returned termperate sure stats
type Temperature struct {
	Day   float64 `json:"day"`
	Min   float64 `json:"min"`
	Max   float64 `json:"max"`
	Night float64 `json:"night"`
	Eve   float64 `json:"eve"`
	Morn  float64 `json:"morn"`
}

// City data for given location
type City struct {
	ID         int         `json:"id"`
	Name       string      `json:"name"`
	Coord      Coordinates `json:"coord"`
	Country    string      `json:"country"`
	Population int         `json:"population"`
	Sys        ForecastSys `json:"sys"`
}

// ForecastWeatherList holds specific query data
type ForecastWeatherList struct {
	Dt       int         `json:"dt"`
	Temp     Temperature `json:"temp"`
	Pressure float64     `json:"pressure"`
	Humidity int         `json:"humidity"`
	Weather  []Weather   `json:"weather"`
	Speed    float64     `json:"speed"`
	Deg      int         `json:"deg"`
	Clouds   int         `json:"clouds"`
	Snow     float64     `json:"snow"`
	Rain     float64     `json:"rain"`
}

// ForecastWeatherData will hold returned data from queries
type ForecastWeatherData struct {
	COD     string                `json:"cod"`
	Message float64               `json:"message"`
	City    City                  `json:"city"`
	Cnt     int                   `json:"cnt"`
	List    []ForecastWeatherList `json:"list"`
	Unit    string
	Lang    string
}

// NewForecast returns a new HistoricalWeatherData pointer with
// the supplied arguments.
func NewForecast(unit, lang string) (*ForecastWeatherData, error) {
	unitChoice := strings.ToUpper(unit)
	langChoice := strings.ToUpper(lang)
	f := &ForecastWeatherData{}
	if ValidDataUnit(unitChoice) {
		f.Unit = unitChoice
	} else {
		return nil, errors.New(unitError)
	}
	if ValidLangCode(langChoice) {
		f.Lang = langChoice
	} else {
		return nil, errors.New(langError)
	}
	return f, nil
}

// DailyByName will provide a forecast for the location given for the
// number of days given.
func (f *ForecastWeatherData) DailyByName(location string, days int) error {
	response, err := http.Get(fmt.Sprintf(forecastBase, "q", location, f.Unit, f.Lang, days))
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if err = json.NewDecoder(response.Body).Decode(&f); err != nil {
		return err
	}
	return nil
}

// DailyByCoordinates will provide a forecast for the coordinates ID give
// for the number of days given.
func (f *ForecastWeatherData) DailyByCoordinates(location *Coordinates, days int) error {
	response, err := http.Get(fmt.Sprintf(fmt.Sprintf(forecastBase, "lat=%f&lon=%f&units=%s"), location.Latitude, location.Longitude, f.Unit, f.Lang, days))
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if err = json.NewDecoder(response.Body).Decode(&f); err != nil {
		return err
	}
	return nil
}

// DailyByID will provide a forecast for the location ID give for the
// number of days given.
func (f *ForecastWeatherData) DailyByID(id, days int) error {
	response, err := http.Get(fmt.Sprintf(forecastBase, "id", strconv.Itoa(id), f.Unit, f.Lang, days))
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if err = json.NewDecoder(response.Body).Decode(&f); err != nil {
		return err
	}
	return nil
}
