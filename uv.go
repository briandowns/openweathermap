package openweathermap

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// CurUV holds the results of a call to the UV API
type CurUV struct {
	Coord []float64 `json:"coord"`
	DT    int64     `json:"dt"`
	Value float64   `json:"value"`
}

// HistUV holds the results of a call to the UV API
type HistUV struct {
	Coord Coordinates `json:"coord"`
	Data  []struct {
		DT    int64   `json:"dt"`
		Value float64 `json:"value"`
	} `json:"data"`
}

// CurrentUV gets the current UV data for the given coordinates
func CurrentUV(coord *Coordinates) (*CurUV, error) {
	response, err := http.Get(fmt.Sprintf("%scurrent?lat=%f&lon=%f&appid=%s", uvURL, coord.Latitude, coord.Longitude, getKey()))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var cuv CurUV

	if err = json.NewDecoder(response.Body).Decode(&cuv); err != nil {
		return nil, err
	}

	return &cuv, nil
}

// HistoricalUV gets the historical
func HistoricalUV(coord *Coordinates, start, end time.Time) (*HistUV, error) {
	response, err := http.Get(fmt.Sprintf("%slist?lat=%f&lon=%f&from=%d&to=%d&appid=%s", uvURL, coord.Latitude, coord.Longitude, start, end, getKey()))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var hist HistUV

	if err = json.NewDecoder(response.Body).Decode(&hist); err != nil {
		return nil, err
	}

	return &hist, nil
}
