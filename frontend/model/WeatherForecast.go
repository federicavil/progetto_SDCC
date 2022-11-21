package model

type WeatherForecasts struct {
	Date        []string  `json:"time"`
	Temperature []float32 `json:"temperature"`
	Cloudy      []int     `json:"cloud_cover"`
	Rainy       []float32 `json:"precipitation"`
	Wind        []float32 `json:"wind_speed"`
	Humidity    []int     `json:"humidity"`
}

type WeatherForecast struct {
	Date        string
	Temperature float32
	Cloudy      int
	Rainy       float32
	Wind        float32
	Humidity    int
}
