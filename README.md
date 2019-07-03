# OpenWeatherMap w/ golang
A lightweight wrapper for the OpenWeatherMap API for use with golang

## Usage

First, create an instance of an OpenWeatherMap struct with your APP ID
```go
package main
import "github.com/ramsgoli/openweathermap"

owm := openweathermap.OpenWeatherMap{API_KEY: os.Getenv("OWM_APP_ID")}
```

### Fetching the current weather

Create an instance of the CurrentWeatherMap struct. 

If fetching the weather from a specific city, pass the city name as an argument to the CurrentWeatherFromCity function
```go
var currentWeather *openweathermap.CurrentWeatherResponse
var err error

currentWeather, err = owm.CurrentWeatherFromCity(city)
```

If fetching the weather from geocoordinates, pass the latitude and longitudes arguments to the CurrentWeatherFromCoordinates function
```go
var currentWeather *openweathermap.CurrentWeatherResponse
var err error

currentWeather, err = owm.CurrentWeatherFromCoordinates(lat, long)
```

If fetching the weather from a city ID, pass the city ID as an argument to the CurrentWeatherFromCityID function
```go
var currentWeather *openweathermap.CurrentWeatherResponse
var err error

currentWeather, err = owm.CurrentWeatherFromCityID(id)
```

If fetching the weather from a zip code, pass the zip code as an argument to the CurrentWeatherFromZip function
```go
var currentWeather *openweathermap.CurrentWeatherResponse
var err error

currentWeather, err = owm.CurrentWeatherFromZip(zip)
```

This function returns a struct, (CurrentWeatherResonse) that matches the fields of the json response from the API
```json
{"coord":{"lon":139,"lat":35},
"sys":{"country":"JP","sunrise":1369769524,"sunset":1369821049},
"weather":[{"id":804,"main":"clouds","description":"overcast clouds","icon":"04n"}],
"main":{"temp":289.5,"humidity":89,"pressure":1013,"temp_min":287.04,"temp_max":292.04},
"wind":{"speed":7.31,"deg":187.002},
"rain":{"3h":0},
"clouds":{"all":92},
"dt":1369824698,
"id":1851632,
"name":"Shuzenji",
"cod":200}
```

To access fields of the struct, use the json object's field name with the first letter capitalized
```go
fmt.Printf("The current temperature in %s is %.2f degrees\n", currentWeather.Name, currentWeather.Main.Temp)
```

