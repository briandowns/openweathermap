package openweathermap

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type ForecastSys struct {
	Population int `json:"population"`
}

type Temperature struct {
	Day   float64 `json:"day"`
	Min   float64 `json:"min"`
	Max   float64 `json:"max"`
	Night float64 `json:"night"`
	Eve   float64 `json:"eve"`
	Morn  float64 `json:"morn"`
}

type City struct {
	ID         int         `json:"id"`
	Name       string      `json:"name"`
	Coord      Coordinates `json:"coord"`
	Country    string      `json:"country"`
	Population int         `json:"population"`
	Sys        ForecastSys `json:"sys"`
}

type ForecastWeatherList struct {
	Dt       int         `json:"dt"`
	Temp     Temperature `json:"temp"`
	Pressure float64     `json:"pressure"`
	Humidity int         `json:"humidity"`
	Weather  []Weather   `json:"weather"`
	Speed    float64     `json:"speed"`
	Deg      int         `json:"deg"`
	Clouds   int         `json:"clouds"`
	Rain     int         `json:"rain"`
}

type ForecastWeatherData struct {
	COD     string                `json:"cod"`
	Message float64               `json:"message"`
	City    City                  `json:"city"`
	Cnt     int                   `json:"cnt"`
	List    []ForecastWeatherList `json:"list"`
	Units   string
}

// NewHistorical returns a new HistoricalWeatherData pointer with the supplied
// arguments.
func NewForecast(unit string) (*ForecastWeatherData, error) {
	unitChoice := strings.ToLower(unit)
	for _, i := range dataUnits {
		if strings.Contains(unitChoice, i) {
			return &ForecastWeatherData{Units: unitChoice}, nil
		}
	}
	return nil, errors.New("ERROR: unit of measure not available")
}

// DailyByName will provide a forecast for the location given for the
// number of days given.
func (f *ForecastWeatherData) DailyByName(location string, days int) {
	response, err := http.Get(fmt.Sprintf(forecastBase, location, f.Units, days))
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(result, &f)
	if err != nil {
		log.Fatalln(err)
	}
}
