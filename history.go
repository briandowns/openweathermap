package openweathermap

import (
	"errors"
	"strings"
)

// HistoricalParameters struct holds the (optional) fields to be
// supplied for historical data requests.
type HistoricalParameters struct {
	Start int64
	End   int64
	Cnt   int
}

// WeatherHistory struct contains aggregate fields from the above
// structs.
type WeatherHistory struct {
	Main    Main      `json:"main"`
	Wind    Wind      `json:"wind"`
	Clouds  Clouds    `json:"clouds"`
	Weather []Weather `json:"weather"`
	Dt      int       `json:"dt"`
}

// HistoricalWeatherData struct is where the JSON is unmarshaled to
// when receiving data for a historical request.
type HistoricalWeatherData struct {
	Message  string           `json:"message"`
	Cod      int              `json:"cod"`
	CityData int              `json:"city_data"`
	CalcTime float64          `json:"calctime"`
	Cnt      int              `json:"cnt"`
	List     []WeatherHistory `json:"list"`
	Units    string
}

// NewHistorical returns a new HistoricalWeatherData pointer with the supplied
// arguments.
func NewHistorical(unit string) (*HistoricalWeatherData, error) {
	unitChoice := strings.ToLower(unit)
	for _, i := range dataUnits {
		if strings.Contains(unitChoice, i) {
			return &HistoricalWeatherData{Units: unitChoice}, nil
		}
	}
	return nil, errors.New("ERROR: unit of measure not available")
}
