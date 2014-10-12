package openweathermap

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Coordinates struct {
	Longitude float64 `json:"lon"`
	Latitude  float64 `json:"lat"`
}

type Sys struct {
	Type    int     `json:"type"`
	Id      int     `json:"id"`
	Message float64 `json:"message"`
	Country string  `json:"country"`
	Sunrise int     `json:"sunrise"`
	Sunset  int     `json:"sunset"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
}

type Weather struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Main struct {
	Temp     float64 `json:"temp"`
	TempMin  float64 `json:"temp_min"`
	TempMax  float64 `json:"temp_max"`
	Pressure int     `json:"pressure"`
	Humidity int     `json:"humidity"`
}

type Clouds struct {
	All int `json:"all"`
}

type WeatherData struct {
	GeoPos  Coordinates `json:"coord"`
	Sys     Sys         `json:"sys"`
	Base    string      `json:"base"`
	Weather []Weather   `json:"weather"`
	Main    Main        `json:"main"`
	Wind    Wind        `json:"wind"`
	Clouds  Clouds      `json:"clouds"`
	Dt      int         `json:"dt"`
	Id      int         `json:"id"`
	Name    string      `json:"name"`
	Cod     int         `json:"cod"`
}

func (w *WeatherData) GetByName(location string) {
	response, err := http.Get(fmt.Sprintf(cityUrl, location))
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading data", err)
	}

	err = json.Unmarshal(result, &w)
	if err != nil {
		log.Fatalln(err)
	}
}

func (w *WeatherData) GetByCoordinates(location *Coordinates) {
	response, err := http.Get(fmt.Sprintf(coordUrl, location.Latitude, location.Longitude))
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading data", err)
	}

	err = json.Unmarshal(result, &w)
	if err != nil {
		log.Fatalln(err)
	}
}

func (w *WeatherData) GetByID() {}

func (w *WeatherData) GetByArea() {}
