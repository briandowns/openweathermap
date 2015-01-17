// Copyright 2015 Brian J. Downs
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//
// weather.go
//
// This application will go out and get the weather for the given
// location and display it in the given data units (fahrenheit,
// celcius, or kelvin.  If the string "here" is provided as an
// argument to the -l flag, the app will try to figure out where
// it's being executed from based on geolocation from IP address.
//
// Examples:
//          go run weather.go --help
//          go run weather.go -w Philadelphia -u f -l en  # fahrenheit, English
//          go run weather.go -w here -u f -l ru          # fahrenheit, Russian
//          go run weather.go -w Dublin -u c -l fi        # celcius, Finnish
//          go run weather.go -w "Las Vegas" -u k -l es   # kelvin, Spanish
package main

import (
	"encoding/json"
	"flag"
	owm "github.com/briandowns/openweathermap" // "owm" for easier use
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

const (
	url = "http://ip-api.com/json"
	// template used for output
	weatherTemplate = `Current weather for {{.Name}}:
    Conditions: {{range .Weather}} {{.Description}} {{end}}
    Now:         {{.Main.Temp}} {{.Unit}}
    High:        {{.Main.TempMax}} {{.Unit}}
    Low:         {{.Main.TempMin}} {{.Unit}}
`
)

// Pointers to hold the contents of the flag args.
var (
	whereFlag = flag.String("w", "", "Location to get weather.  If location has a space, wrap the location in double quotes.")
	unitFlag  = flag.String("u", "", "Unit of measure to display temps in")
	langFlag  = flag.String("l", "", "Language to display temps in")
)

// Data will hold the result of the query to get the IP
// address of the caller.
type Data struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	ISP         string  `json:"isp"`
	ORG         string  `json:"org"`
	AS          string  `json:"as"`
	Message     string  `json:"message"`
	Query       string  `json:"query"`
}

// getLocation will get the location details for where this
// application has been run from.
func getLocation() *Data {
	response, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	r := &Data{}
	err = json.Unmarshal(result, &r)
	if err != nil {
		log.Fatalln(err)
	}
	return r
}

// getCurrent gets the current weather for the provided
// location in the units provided.
func getCurrent(l, u, lang string) *owm.CurrentWeatherData {
	w, err := owm.NewCurrent(u, lang)
	if err != nil {
		log.Fatalln(err)
	}
	w.CurrentByName(l)
	return w
}

func main() {
	flag.Parse()

	// If there's any funkiness with cli args, tuck and roll...
	if len(*whereFlag) <= 1 || len(*unitFlag) != 1 || len(*langFlag) != 2 {
		flag.Usage()
		os.Exit(1)
	}

	// Process request for location of "here"
	if strings.ToLower(*whereFlag) == "here" {
		w := getCurrent(getLocation().City, *unitFlag, *langFlag)
		tmpl, err := template.New("weather").Parse(weatherTemplate)
		if err != nil {
			log.Fatalln(err)
		}

		// Render the template and display
		err = tmpl.Execute(os.Stdout, w)
		if err != nil {
			log.Fatalln(err)
		}
		os.Exit(0)
	}

	// Process request for the given location
	w := getCurrent(*whereFlag, *unitFlag, *langFlag)
	tmpl, err := template.New("weather").Parse(weatherTemplate)
	if err != nil {
		log.Fatalln(err)
	}

	// Render the template and display
	err = tmpl.Execute(os.Stdout, w)
	if err != nil {
		log.Fatalln(err)
	}
	os.Exit(0)
}
