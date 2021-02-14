package openweathermap

import (
	"fmt"
)

const (
	oneCallBaseURL = "https://api.openweathermap.org/data/2.5/onecall?%s"
)

// OneTimeCurrent
type OneTimeCurrent struct {
	Dt         int64     `json:"dt"`
	Sunrise    int64     `json:"sunrise"`
	Sunset     int64     `json:"sunset"`
	Temp       float64   `json:"temp"`
	FeelsLike  float64   `json:"feels_like"`
	Pressure   float64   `json:"pressure"`
	Humidity   int64     `json:"humidity"`
	DewPoint   float64   `json:"dew_point"`
	UVI        int64     `json:"uvi"`
	Clouds     int64     `json:"clouds"`
	Visibility int64     `json:"visibility"`
	WindSpeed  float64   `json:"wind_speed"`
	WindDeg    float64   `json:"wind_deg"`
	Weather    []Weather `json:"weather"`
}

// OneTimeMinutely
type OneTimeMinutely struct {
	Dt            int64 `json:"dt"`
	Precipitation int64 `json:"precipitation"`
}

// OneTimeHourly
type OneTimeHourly struct {
	Dt         int64     `json:"dt"`
	Temp       float64   `json:"temp"`
	FeelsLike  float64   `json:"feels_like"`
	Pressure   float64   `json:"pressure"`
	Humidity   int64     `json:"humidity"`
	DewPoint   float64   `json:"dew_point"`
	UVI        int64     `json:"uvi"`
	Clouds     int64     `json:"clouds"`
	Visibility int64     `json:"visibility"`
	WindSpeed  float64   `json:"wind_speed"`
	WindDeg    float64   `json:"wind_deg"`
	Weather    []Weather `json:"weather"`
	Pop        float64   `json:"pop"`
}

// TempFullDay
type TempFullDay struct {
	Day   float64 `json:"day"`
	Min   float64 `json:"Min"`
	Max   float64 `json:"max"`
	Night float64 `json:"night"`
	Eve   float64 `json:"eve"`
	Morn  float64 `json:"morn"`
}

// FeelsLikeFullDay
type FeelsLikeFullDay struct {
	Day   float64 `json:"day"`
	Night float64 `json:"night"`
	Eve   float64 `json:"eve"`
	Morn  float64 `json:"morn"`
}

// OneTimeDaily
type OneTimeDaily struct {
	Dt        int64             `json:"dt"`
	Sunrise   int64             `json:"sunrise"`
	Sunset    int64             `json:"sunset"`
	Temp      *TempFullDay      `json:"tmep"`
	FeelsLike *FeelsLikeFullDay `json:"feels_like"`
	Pressure  float64           `json:"pressure"`
	Humidity  int64             `json:"humidity"`
	DewPoint  float64           `json:"dew_point"`
	WindSpeed float64           `json:"wind_speed"`
	WindDeg   float64           `json:"wind_deg"`
	Weather   []Weather         `json:"weather"`
	Clouds    int64             `json:"clouds"`
	Pop       float64           `json:"pop"`
	UVI       int64             `json:"uvi"`
}

// Alert
type Alert struct {
	SenderName  string `json:"sender_name"`
	Event       string `json:"event"`
	Start       int64  `json:"start"`
	End         int64  `json:"end"`
	Description string `json:"description"`
}

// OneCallData
type OneCallData struct {
	Longitude      float64           `json:"lon"`
	Latitude       float64           `json:"lat"`
	Timezone       string            `json:"timezone"`
	TimezoneOffset int64             `json:"timezone_offset"`
	Current        *OneTimeCurrent   `json:"current"`
	Weather        *Weather          `json:"weather"`
	Minutely       []OneTimeMinutely `json:"minutely"`
	Hourly         []OneTimeHourly   `json:"hourly"`
	Daily          []OneTimeDaily    `json:"daily"`
	Alerts         []Alert           `json:"alerts"`
}

// OneCallCurrentAndForecast
func (o *OWM) OneCallCurrentAndForecast(location *Coordinates) (*OneCallData, error) {
	base := fmt.Sprintf(oneCallBaseURL, "lat=%s&lon=%s&exclude=%s&lang=%s&appid=%s")
	url := fmt.Sprintf(base, location.Latitude, location.Longitude, o.unit, o.lang, o.apiKey)

	var otc OneCallData
	if err := o.call(url, &otc); err != nil {
		return nil, err
	}

	return &otc, nil
}

// OneCallHistorical
func (o *OWM) OneCallHistorical(location *Coordinates, dt int64) (*OneCallData, error) {
	base := fmt.Sprintf(oneCallBaseURL, "lat=%s&lon=%s&exclude=%s&lang=%s&appid=%s")
	url := fmt.Sprintf(base, location.Latitude, location.Longitude, o.unit, o.lang, o.apiKey)

	var otc OneCallData
	if err := o.call(url, &otc); err != nil {
		return nil, err
	}

	return &otc, nil
}
