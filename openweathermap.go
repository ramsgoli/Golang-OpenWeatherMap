package openweathermap

/*
Define API response fields
 */
type OpenWeatherMap struct {
	API_KEY string
}

type coord struct {
	lon int
	lat int
}

type weather struct {
	id int
	main string
	description string
}

type wind struct {
	speed float64
	deg float64
}

type clouds struct {
	all int
}

type rain struct {
	threehr int `json:"3h"`
}

type visibility string

type dt int

type id int

type name string

type main struct {
	temp float64
	pressure int
	humidity int
	temp_min float64
	temp_max float64
}

type forecastResponse struct {
	city struct {
		id int
		name string
	}
	coord

	country string
	cnt int
	list [] struct {
		dt string
		main struct {
			temp float64
			temp_min float64
			temp_max float64
			humidity float64
		}
		weather []weather

		clouds struct {
			all int
		}
		wind struct {
			speed float64
			deg float64
		}
	}
}

func (owm *OpenWeatherMap) currentWeather() {

}
