package openweathermap

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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
	Units   string
}

// NewForecast returns a new HistoricalWeatherData pointer with
// the supplied arguments.
func NewForecast(unit string) (*ForecastWeatherData, error) {
	unitChoice := strings.ToLower(unit)
	if ValidDataUnit(unitChoice) {
		return &ForecastWeatherData{Units: unitChoice}, nil
	}
	return nil, errors.New("unit of measure not available")
}

// DailyByName will provide a forecast for the location given for the
// number of days given.
func (f *ForecastWeatherData) DailyByName(location string, days int) error {
	response, err := http.Get(fmt.Sprintf(forecastBase, "q", location, f.Units, days))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(result, &f)
	if err != nil {
		return err
	}
	return nil
}

// DailyByCoordinates will provide a forecast for the coordinates ID give
// for the number of days given.
func (f *ForecastWeatherData) DailyByCoordinates(location *Coordinates, days int) error {
	response, err := http.Get(fmt.Sprintf(fmt.Sprintf(forecastBase, "lat=%f&lon=%f&units=%s"), location.Latitude, location.Longitude, f.Units, days))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(result, &f)
	if err != nil {
		return err
	}
	return nil
}

// DailyByID will provide a forecast for the location ID give for the
// number of days given.
func (f *ForecastWeatherData) DailyByID(id, days int) error {
	response, err := http.Get(fmt.Sprintf(forecastBase, "id", strconv.Itoa(id), f.Units, days))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(result, &f)
	if err != nil {
		return err
	}
	return nil
}
