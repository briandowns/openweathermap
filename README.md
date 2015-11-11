# OpenWeatherMap Go API

[![GoDoc](https://godoc.org/github.com/briandowns/openweathermap?status.svg)](https://godoc.org/github.com/briandowns/openweathermap) [![Build Status](https://travis-ci.org/briandowns/openweathermap.svg?branch=master)](https://travis-ci.org/briandowns/openweathermap)

Go (golang) package for use with openweathermap.org's API.

For more detail about the library and its features, reference your local godoc once installed.

[Website](https://briandowns.github.io/openweathermap)!

To use the OpenweatherMap API, you need to obtain an API key.  Sign up [here](http://home.openweathermap.org/users/sign_up).  Once you have your key, create an environment variable called `OWM_API_KEY`.  Start coding!

[Slack Channel](https://openweathermapgolang.slack.com/messages/general)

Contributions welcome!

## Features

### Current Weather Conditions

- By City
- By City,St (State)
- By City,Co (Country)
- By City ID
- By Zip,Co (Country)
- By Longitude and Latitude

## Forecast

Get the weather conditions for a given number of days.

- By City
- By City,St (State)
- By City,Co (Country)
- By City ID
- By Longitude and Latitude

### Access to Condition Codes and Icons

Gain access to OpenWeatherMap icons and condition codes.

- Thunderstorms
- Drizzle
- Rain
- Snow
- Atmosphere
- Clouds
- Extreme
- Additional

### Data Available in Multiple Measurement Systems

- Fahrenheit (OpenWeatherMap API - imperial)
- Celcius (OpenWeatherMap API - metric)
- Kelvin (OpenWeatherMap API - internal)

## Historical Conditions

- ...still in the works...

## Supported Languages

English - en, Russian - ru, Italian - it, Spanish - es (or sp), Ukrainian - uk (or ua), German - de, Portuguese - pt, Romanian - ro, Polish - pl, Finnish - fi, Dutch - nl, French - fr, Bulgarian - bg, Swedish - sv (or se), Chinese Traditional - zh_tw, Chinese Simplified - zh (or zh_cn), Turkish - tr, Croatian - hr, Catalan - ca

## Installation

```bash
go get github.com/briandowns/openweathermap
```

## Examples

There are a few full examples in the examples directory that can be referenced.  1 is a command line application and 1 is a simple web application.

```Go
package main

import (
    "log"
    "fmt"

	// Shortening the import reference name seems to make it a bit easier
    owm "github.com/briandowns/openweathermap"
)

func main() {
    w, err := owm.NewCurrent("F", "ru") // fahrenheit (imperial) with Russian output
    if err != nil {
        log.Fatalln(err)
    }

    w.CurrentByName("Phoenix")
    fmt.Println(w)
}
```

### Current Conditions by location name

```Go
func main() {
    w, err := owm.NewCurrent("K", "EN") // (internal - OpenWeatherMap reference for kelvin) with English output
    if err != nil {
        log.Fatalln(err)
    }

    w.CurrentByName("Phoenix,AZ")
    fmt.Println(w)
}
```

### Forecast Conditions in imperial (fahrenheit) by coordinates

```Go
func main() {
    w, err := owm.NewForecast("F", "FI")
    if err != nil {
        log.Fatalln(err)
    }

    w.DailyByCoordinates(
    		&Coordinates{
    			Longitude: -112.07,
    			Latitude: 33.45,
    		},
    )
    fmt.Println(w)
}
```

### Current conditions in metric (celsius) by location ID

```Go
func main() {
    w, err := owm.NewCurrent("C", "PL")
    if err != nil {
        log.Fatalln(err)
    }

    w.CurrentByID(2172797)
    fmt.Println(w)
}
```

### Current conditions by zip code. 2 character country code required

```Go
func main() {
	w, err := owm.NewCurrent("F", "US")
	if err != nil {
		log.Fatalln(err)
	}

	w.CurrentByZip(19125, "US")
	fmt.Println(w)
}
```
