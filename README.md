# OpenWeatherMap Go SDK

[![GoDoc](https://godoc.org/github.com/briandowns/openweathermap?status.svg)](https://godoc.org/github.com/briandowns/openweathermap) [![Build Status](https://travis-ci.org/briandowns/openweathermap.svg?branch=master)](https://travis-ci.org/briandowns/openweathermap)

For more detail about the library and its features, reference your local godoc once installed.

## Features 

- Conditions by City
- Conditions by City,St (State)
- Conditions by City,Co (Country)
- Conditions by City ID
- Conditions by Longitude and Latitude

## Installation

```bash
go get github.com/briandowns/openweathermap
```

## Examples

```Go
package main

import github.com/briandowns/openweathermap

func main() {
    w, err := New("imperial")
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
    w, err := New("metric")
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
    w, err := New("imperial")
    if err != nil {
        log.Fatalln(err)
    }
    
    c := &Coordinates{
    		Longitude: -112.07,
    		Latitude: 33.45,
    }
    
    w.GetByLocation(c)
    fmt.Println(w)
}
```
```bash
```

```Go
func main() {
    w, err := New("metric")
    if err != nil {
        log.Fatalln(err)
    }
    
    w.GetByID(2172797)
    fmt.Println(w)
}
```
```bash
```