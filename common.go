package openweathermap

const (
	baseUrl      string = "http://api.openweathermap.org/data/2.5/weather?%s"
	iconUrl      string = "http://openweathermap.org/img/w/%s"
	stationUrl   string = "http://api.openweathermap.org/data/2.5/station?id=%d"
	forecastBase string = "http://api.openweathermap.org/data/2.5/forecast/daily?%s=%s&mode=json&units=%s&cnt=%d"
)

var (
	dataUnits = [3]string{"metric", "imperial", "internal"}
)

// APIError returned on failed API calls.
type APIError struct {
	Message string `json:"message"`
	COD     string `json:"cod"`
}

// Coordinates struct holds longitude and latitude data
// in returned JSON or as parameter data for requests
// using longitude and latitude.
type Coordinates struct {
	Longitude float64 `json:"lon"`
	Latitude  float64 `json:"lat"`
}

// Sys struct contains general information about the request and the
// surrounding area for where the request was made.
type Sys struct {
	Type    int     `json:"type"`
	Id      int     `json:"id"`
	Message float64 `json:"message"`
	Country string  `json:"country"`
	Sunrise int     `json:"sunrise"`
	Sunset  int     `json:"sunset"`
}

// Wind struct contains the speed and degree of the wind.
type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
}

// Weather struct holds high-level basic info on the returned
// data.
type Weather struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

// Main struct contains the meat and potatos of the request.
type Main struct {
	Temp     float64 `json:"temp"`
	TempMin  float64 `json:"temp_min"`
	TempMax  float64 `json:"temp_max"`
	Pressure int     `json:"pressure"`
	Humidity int     `json:"humidity"`
}

// Clouds struct holds data regarding cloud cover.
type Clouds struct {
	All int `json:"all"`
}
