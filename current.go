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

func New(unit string) (*WeatherData, error) {
	unitChoice := strings.ToLower(unit)
	for _, i := range dataUnits {
		if strings.Contains(unitChoice, i) {
			return &WeatherData{Units: unitChoice}, nil
		}
	}
	return nil, errors.New("ERROR: unit of measure not available")
}

func (w *WeatherData) GetByName(location string) {
	response, err := http.Get(fmt.Sprintf(fmt.Sprintf(baseUrl, "q=%s&units=%s"), location, w.Units))
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(result, &w)
	if err != nil {
		log.Fatalln(err)
	}
}

func (w *WeatherData) GetByCoordinates(location *Coordinates) {
	response, err := http.Get(fmt.Sprintf(fmt.Sprintf(baseUrl, "lat=%f&lon=%f&units=%s"), location.Latitude, location.Longitude, w.Units))
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(result, &w)
	if err != nil {
		log.Fatalln(err)
	}
}

func (w *WeatherData) GetByID(id int) {
	response, err := http.Get(fmt.Sprintf(fmt.Sprintf(baseUrl, "id=%d&units=%s"), id, w.Units))
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(result, &w)
	if err != nil {
		log.Fatalln(err)
	}
}

func (w *WeatherData) GetByArea() {}
