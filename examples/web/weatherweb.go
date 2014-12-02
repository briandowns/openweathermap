// Example of creating a web based application purely using
// the net/http package to display weather information and
// Twitter Bootstrap so it doesn't look like it's '92.
//
// To start the app, run:
//    go run weatherweb.go
//
// Accessible via:  http://localhost:8888/here
package main

import (
	"encoding/json"
	owm "github.com/briandowns/openweathermap"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	url = "http://ip-api.com/json"
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
		log.Fatal(err)
	}
	defer response.Body.Close()

	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	r := &Data{}
	err = json.Unmarshal(result, &r)
	if err != nil {
		log.Fatal(err)
	}
	return r
}

// getCurrent gets the current weather for the provided location in
// the units provided.
func getCurrent(l, u string) *owm.CurrentWeatherData {
	w, err := owm.NewCurrent(u)
	if err != nil {
		log.Fatal(err)
	}
	w.CurrentByName(l)
	return w
}

// hereHandler will take are of requests coming in for the "/here" route.
func hereHandler(w http.ResponseWriter, r *http.Request) {
	wd := getCurrent(getLocation().City, "imperial")
	wd.Units = owm.DataUnits[wd.Units]

	t, err := template.ParseFiles("templates/here.html")
	if err != nil {
		log.Fatal(err)
	}
	// We're doing naughty things below... Ignoring icon file size and possible errors.
	_, _ = owm.RetrieveIcon("static/img", wd.Weather[0].Icon+".png")
	t.Execute(w, wd)
}

// Run the app
func main() {
	http.HandleFunc("/here", hereHandler)
	// Make sure we can serve our icon files once retrieved
	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	http.ListenAndServe(":8888", nil)
}
