package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type apiConfiDate struct {
	OpenWeatherMapApiKey string `json:"OpenWeatherMapApiKey"`
}

type weatherData struct {
	Name string `json:"name"`
	Main        `json:"main"`
}

type Main struct {
	Kelvin float64 `json:"temp"`
	Celcius float64 `json:"celcius"`
}


func loadApiConfig(filename string) (apiConfiDate, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
	return apiConfiDate{}, err
	}
	var c apiConfiDate
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return apiConfiDate{}, err
	}
	return c, nil
}

func main(){

}