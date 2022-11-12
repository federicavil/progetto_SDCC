package model

type WeatherForecasts struct {
	Date        []string `json:"time"`
	Temperature []int    `json:"temperature"`
	Cloudy      []int    `json:"cloud_cover"`
	Rainy       []int    `json:"precipitation"`
	Wind        []int    `json:"wind_speed"`
	Humidity    []int    `json:"humidity"`
}

type WeatherForecast struct {
	Date        string
	Temperature int
	Cloudy      int
	Rainy       int
	Wind        int
	Humidity    int
}
