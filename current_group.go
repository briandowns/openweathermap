package openweathermap

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// maximum count of several identifiers
// used in CurrentByIDs
const maxCityIDs = 20

// CurrentWeatherGroup struct contains list of the CurrentWeatherData
// structs for JSON to be unmarshaled into.
type CurrentWeatherGroup struct {
	Count int                   `json:"count"`
	List  []*CurrentWeatherData `json:"list,omitempty"`

	Unit string
	Lang string
	Key  string

	*Settings
}

// NewCurrentGroup returns a new CurrentWeatherGroup pointer with the supplied parameters
func NewCurrentGroup(unit, lang, key string, options ...Option) (*CurrentWeatherGroup, error) {
	unitChoice := strings.ToUpper(unit)
	langChoice := strings.ToUpper(lang)

	g := &CurrentWeatherGroup{
		Settings: NewSettings(),
	}

	if ValidDataUnit(unitChoice) {
		g.Unit = DataUnits[unitChoice]
	} else {
		return nil, errUnitUnavailable
	}

	if ValidLangCode(langChoice) {
		g.Lang = langChoice
	} else {
		return nil, errLangUnavailable
	}
	var err error
	g.Key, err = setKey(key)
	if err != nil {
		return nil, err
	}

	if err := setOptions(g.Settings, options); err != nil {
		return nil, err
	}
	return g, nil
}

// CurrentByIDs will provide the current weather as a list
// by the specified location identifiers
func (g *CurrentWeatherGroup) CurrentByIDs(ids ...int) error {
	n := len(ids)
	if n > maxCityIDs {
		return errCountOfCityIDs
	}

	strIDs := make([]string, n)
	for i, id := range ids {
		strIDs[i] = strconv.Itoa(id)
	}

	id := strings.Join(strIDs, ",")
	uri := fmt.Sprintf(groupURL, "appid=%s&id=%s&units=%s&lang=%s")

	response, err := g.client.Get(fmt.Sprintf(uri, g.Key, id, g.Unit, g.Lang))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		return errInvalidKey
	}

	if err = json.NewDecoder(response.Body).Decode(&g); err != nil {
		return err
	}

	for _, w := range g.List {
		w.Settings = g.Settings
		w.Unit = g.Unit
		w.Lang = g.Lang
		w.Key = g.Key
	}

	return nil
}
