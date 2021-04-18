# OpenWeatherMap Go API

[![GoDoc](https://godoc.org/github.com/briandowns/openweathermap?status.svg)](https://godoc.org/github.com/briandowns/openweathermap) [![Build Status](https://travis-ci.org/briandowns/openweathermap.svg?branch=master)](https://travis-ci.org/briandowns/openweathermap) [![Coverage Status](https://coveralls.io/repos/github/briandowns/openweathermap/badge.svg?branch=master)](https://coveralls.io/github/briandowns/openweathermap?branch=master)

Go (golang) package for use with openweathermap.org's HTTP API.

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
- Celsius (OpenWeatherMap API - metric)
- Kelvin (OpenWeatherMap API - internal)

### UV Index Data

- Current
- Historical

### Pollution Data

- Current

## Historical Conditions

- By Name
- By ID
- By Coordinates

## Supported Languages

- English - en
- Russian - ru
- Italian - it
- Spanish - es (or sp)
- Ukrainian - uk (or ua)
- German - de
- Portuguese - pt
- Romanian - ro
- Polish - pl
- Finnish - fi
- Dutch - nl
- French - fr
- Bulgarian - bg
- Swedish - sv (or se)
- Chinese Traditional - zh_tw
- Chinese Simplified - zh (or zh_cn)
- Turkish - tr
- Croatian - hr
- Catalan - ca

## Installation

```bash
go get github.com/briandowns/openweathermap
```

## Examples

Full, simple example.

```Go
package main

import (
	"log"
	"fmt"
	"os"

	"github.com/briandowns/openweathermap"
)

func main() {
    opts := openweathermap.Opts{
		Lang: "EN",
		Unit: "F",
		Client: &http.Client{
			Timeout: time.Second * 5,
		},
	}
	owm, err := openweathermap.New(&opts)
	if err != nil {
		log.Fatalln(err)
	}

	cbn, err := owm.CurrentByName("Philadelphia")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%#v\n", cbn)
}
```

### Forecast Conditions in imperial (fahrenheit) by coordinates

```Go
fdfbc, err := owm.FiveDayForecastByCoordinates(&openweathermap.Coordinates{Longitude: -75.1638, Latitude: 39.9523}, 10)
if err != nil {
    log.Fatalln(err)
}
fmt.Printf("%#v\n", fdfbc)
```

### Current conditions in metric (celsius) by location ID

```Go
owm.Unit = "C"
cbi, err := owm.CurrentByID(4560349)
if err != nil {
    log.Fatalln(err)
}
fmt.Printf("%#v\n", cbi)
```

### Current conditions by zip code. 2 character country code required

```Go
cbz, err := owm.CurrentByZip("19127", "")
if err != nil {
    log.Fatalln(err)
}
fmt.Printf("%#v\n", cbz)
```

### History by Name

```Go
hbn, err := owm.HistoryByName("Philadelphia", &openweathermap.HistoricalParameters{
    Start: 1369728000,
    End:   1369789200,
    Cnt:   4,
})
if err != nil {
    log.Fatalln(err)
}
fmt.Printf("%#v\n", hbn)
```

### Current UV conditions

```Go
uv, err := owm.UVCurrent(&openweathermap.Coordinates{
    Latitude:  39.9523,
    Longitude: -75.1638,
})
if err != nil {
    log.Fatalln(err)
}
fmt.Printf("%#v\n", uv)
```

### Pollution Information

```Go
p, err := owm.PollutionByParams(&openweathermap.PollutionParameters{
    Location: openweathermap.Coordinates{
        Latitude:  39.9523,
        Longitude: -75.1638,
    },
    Datetime: "2006-01-02T15:04:05-0700",
})
if err != nil {
    log.Fatalln(err)
}
fmt.Printf("%#v\n", p)
```