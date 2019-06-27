package openweathermap

import (
	"errors"
	"fmt"
	"net/http"
	"time"
	"io/ioutil"
	"encoding/json"
)

/*
Define API response fields
 */
type OpenWeatherMap struct {
	API_KEY string
}

type City struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Weather struct {
	Id int `json:"id"`
	Main string `json:"main"`
	Description string `json:"description"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg float64 `json:"deg"`
}

type Clouds struct {
	All int `json:"all"`
}

type Rain struct {
	Threehr int `json:"3h"`
}

type Main struct {
	Temp float64 `json:"temp"`
	Pressure int `json:"pressure"`
	Humidity int `json:"humidity"`
	Temp_min float64 `json:"temp_min"`
	Temp_max float64 `json:"temp_max"`
}

/*
Define API response objects (compose of the above fields)
 */

type CurrentWeatherResponse struct {
	Coord `json:"coord"`
	Weather []Weather `json:"weather"`
	Main `json:"main"`
	Wind `json:"wind"`
	Rain `json:"rain"`
	Clouds `json:"clouds"`
	Dt int `json:"dt"`
	Id int `json:"id"`
	Name string `json:"name"`
}

type ForecastResponse struct {
	City `json:"city"`
	Coord `json:"coord"`
	Country string `json:"country"`
	List [] struct {
		Dt int `json:"dt"`
		Main `json:"main"`
		Weather `json:"weather"`
		Clouds `json:"clouds"`
		Wind `json:"wind"`
	} `json:"list"`
}

const (
	API_URL string = "api.openweathermap.org"
)

func makeApiRequest(url string) ([]byte, error) {
	// Build an http client so we can have control over timeout
	client := &http.Client{
		Timeout: time.Second * 2,
	}

	res, getErr := client.Get(url)
	if getErr != nil {
		return nil, getErr
	}

	// defer the closing of the res body
	defer res.Body.Close()

	// read the http response body into a byte stream
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return nil, readErr
	}

	return body, nil
}

func (owm *OpenWeatherMap) CurrentWeatherFromCity(city string) (*CurrentWeatherResponse, error) {
	if (owm.API_KEY == "") {
		// No API keys present, return error
		return nil, errors.New("No API keys present")
	}
	url := fmt.Sprintf("http://%s/data/2.5/weather?q=%s&units=imperial&APPID=%s", API_URL, city, owm.API_KEY)

	body, err := makeApiRequest(url)
	if (err != nil) {
		return nil, err
	}
	var cwr CurrentWeatherResponse

	// unmarshal the byte stream into a Go data type
	jsonErr := json.Unmarshal(body, &cwr)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return &cwr, nil
}

func (owm *OpenWeatherMap) CurrentWeatherFromCoordinates(lat, long float64) (*CurrentWeatherResponse, error) {
	if (owm.API_KEY == "") {
		// No API keys present, return error
		return nil, errors.New("No API keys present")
	}

	url := fmt.Sprintf("http://%s/data/2.5/weather?lat=%f&lon=%f&units=imperial&APPID=%s", API_URL, lat, long, owm.API_KEY)

	body, err := makeApiRequest(url)
	if (err != nil) {
		return nil, err
	}

	var cwr CurrentWeatherResponse

	// unmarshal the byte stream into a Go data type
	jsonErr := json.Unmarshal(body, &cwr)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return &cwr, nil
}

func (owm *OpenWeatherMap) CurrentWeatherFromZip(zip int) (*CurrentWeatherResponse, error) {
	if owm.API_KEY == "" {
		// No API keys present, return error
		return nil, errors.New("No API keys present")
	}
	url := fmt.Sprintf("http://%s/data/2.5/weather?zip=%d&units=imperial&APPID=%s", API_URL, zip, owm.API_KEY)

	body, err := makeApiRequest(url)
	if err != nil {
		return nil, err
	}
	var cwr CurrentWeatherResponse

	// unmarshal the byte stream into a Go data type
	jsonErr := json.Unmarshal(body, &cwr)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return &cwr, nil
}

