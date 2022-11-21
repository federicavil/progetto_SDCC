package model

type WeatherForecasts struct {
	Date        []string  `json:"time"`
	Temperature []float32 `json:"temperature"`
	Cloudy      []int     `json:"cloud_cover"`
	Rainy       []int     `json:"precipitation"`
	Wind        []float32 `json:"wind_speed"`
	Humidity    []int     `json:"humidity"`
}

type WeatherForecast struct {
	Date        string
	Temperature float32
	Cloudy      int
	Rainy       int
	Wind        float32
	Humidity    int
}
