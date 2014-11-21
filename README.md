# OpenWeatherMap Go API

[![GoDoc](https://godoc.org/github.com/briandowns/openweathermap?status.svg)](https://godoc.org/github.com/briandowns/openweathermap) [![Build Status](https://travis-ci.org/briandowns/openweathermap.svg?branch=master)](https://travis-ci.org/briandowns/openweathermap)

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

Full example in the examples directory.

```Go
package main

import (
    "log"
    "fmt"
    "github.com/briandowns/openweathermap"
)

func main() {
    w, err := openweathermap.NewCurrent("imperial")
    if err != nil {
        log.Fatalln(err)
    }

    w.CurrentByName("Phoenix")
    fmt.Println(w)
}
```

```Go
func main() {
    w, err := openweathermap.NewCurrent("metric")
    if err != nil {
        log.Fatalln(err)
    }

    w.CurrentByName("Phoenix,AZ")
    fmt.Println(w)
}
```

```Go
func main() {
    w, err := openweathermap.NewCurrent("imperial")
    if err != nil {
        log.Fatalln(err)
    }

    w.CurrentByCoordinates(
    		&Coordinates{
    			Longitude: -112.07,
    			Latitude: 33.45,
    		},
    )
    fmt.Println(w)
}
```

```Go
func main() {
    w, err := openweathermap.NewCurrent("metric")
    if err != nil {
        log.Fatalln(err)
    }

    w.CurrentByID(2172797)
    fmt.Println(w)
}
```
