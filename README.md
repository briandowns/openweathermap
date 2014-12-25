# OpenWeatherMap Go API

[![GoDoc](https://godoc.org/github.com/briandowns/openweathermap?status.svg)](https://godoc.org/github.com/briandowns/openweathermap) [![Build Status](https://travis-ci.org/briandowns/openweathermap.svg?branch=master)](https://travis-ci.org/briandowns/openweathermap)

Go (golang) package for use with openweathermap.org's API.

For more detail about the library and its features, reference your local godoc once installed.

Contributions welcome!

## Features

### Current Weather Conditions

- By City
- By City,St (State)
- By City,Co (Country)
- By City ID
- By Longitude and Latitude

## Forecast

For a given number of days.

- By City
- By City,St (State)
- By City,Co (Country)
- By City ID
- By Longitude and Latitude

### Access to Condition Codes and Icons

- Thunderstorms
- Drizzle
- Rain
- Snow
- Atmosphere
- Clouds
- Extreme
- Additional

## Historical Conditions

- ...still in the works...

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
    w, err := owm.NewCurrent("imperial")
    if err != nil {
        log.Fatalln(err)
    }

    w.CurrentByName("Phoenix")
    fmt.Println(w)
}
```

### Current Conditions in metric by location name

```Go
func main() {
    w, err := owm.NewCurrent("metric")
    if err != nil {
        log.Fatalln(err)
    }

    w.CurrentByName("Phoenix,AZ")
    fmt.Println(w)
}
```

### Forecast Conditions in imperial by coordinates

```Go
func main() {
    w, err := owm.NewForecast("imperial")
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

### Current conditions in metric by location ID

```Go
func main() {
    w, err := owm.NewCurrent("metric")
    if err != nil {
        log.Fatalln(err)
    }

    w.CurrentByID(2172797)
    fmt.Println(w)
}
```
