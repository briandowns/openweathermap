# OpenWeatherMap Go SDK

[![GoDoc](https://godoc.org/github.com/briandowns/openweathermap?status.svg)](https://godoc.org/github.com/briandowns/openweathermap) [![Build Status](https://travis-ci.org/briandowns/openweathermap.svg?branch=master)](https://travis-ci.org/briandowns/openweathermap)

For more detail about the library and its features, reference your local godoc once installed.

## Features 

### Current Weather Conditions

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

## Installation

```bash
go get github.com/briandowns/openweathermap
```

## Examples

```Go
package main

import github.com/briandowns/openweathermap

func main() {
    w, err := openweathermap.New("imperial")
    if err != nil {
        log.Fatalln(err)
    }
    
    w.GetByName("Phoenix")
    fmt.Println(w)
}
```
```bash
```

```Go
func main() {
    w, err := openweathermap.New("metric")
    if err != nil {
        log.Fatalln(err)
    }
    
    w.GetByName("Phoenix,AZ")
    fmt.Println(w)
}
```
```bash
```

```Go
func main() {
    w, err := openweathermap.New("imperial")
    if err != nil {
        log.Fatalln(err)
    }
    
    w.GetByCoordinates(
    		&Coordinates{
    			Longitude: -112.07,
    			Latitude: 33.45,
    		},
    )
    fmt.Println(w)
}
```
```bash
```

```Go
func main() {
    w, err := openweathermap.New("metric")
    if err != nil {
        log.Fatalln(err)
    }
    
    w.GetByID(2172797)
    fmt.Println(w)
}
```
```bash
```