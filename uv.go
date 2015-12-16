package openweathermap

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// UVResult
type UV struct {
	Coord []float64 `json:"coord"`
	Data  []struct {
		DT    int64   `json:"dt"`
		Value float64 `json:"value"`
	} `json:"data,omitempty"`
	DT    int64   `json:"dt,omitempty"`
	Value float64 `json:"value,omitempty"`
}

// NewUV
func NewUV() *UV {
	return &UV{}
}

// Current gets the current UV data for the given coordinates
func (u *UV) Current(coord *Coordinates) error {
	response, err := http.Get(fmt.Sprintf("%scurrent?lat=%f&lon=%f&appid=%s", uvURL, coord.Latitude, coord.Longitude, getKey()))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if err = json.NewDecoder(response.Body).Decode(&u); err != nil {
		return err
	}

	return nil
}

// Historical gets the historical UV data for the coordinates and times
func (u *UV) Historical(coord *Coordinates, start, end time.Time) error {
	response, err := http.Get(fmt.Sprintf("%slist?lat=%f&lon=%f&from=%d&to=%d&appid=%s", uvURL, coord.Latitude, coord.Longitude, start, end, getKey()))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if err = json.NewDecoder(response.Body).Decode(&u); err != nil {
		return err
	}

	return nil
}

// UVInformation provides information on the given UV data which includes the severity
// and "Recommended protection"
func (u *UV) UVInformation() {

}
