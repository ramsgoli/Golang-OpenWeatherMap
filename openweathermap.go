package openweathermap

type OpenWeatherMap struct {
	API_KEY string
}

type forecastResponse struct {
	city struct {
		id int
		name string
	}
	coord struct {
		long float64
		lat float64
	}
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
		weather [] struct {
			id int
			main string
			description string
		}
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
