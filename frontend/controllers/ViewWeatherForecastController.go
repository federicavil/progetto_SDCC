package controllers

import (
	"encoding/json"
	"fmt"
	"frontend/conf"
	"frontend/model"
	"github.com/astaxie/beego"
	"github.com/beego/beego/v2/client/httplib"
)

type ViewWeatherForecastController struct {
	beego.Controller
}

func (this *ViewWeatherForecastController) Prepare() {
	this.TplName = "viewWeatherForecast.html"
}

func (this *ViewWeatherForecastController) Get() {
	path := this.StartSession().Get("selectedPath").(model.MountainPath)
	pathJson, _ := json.Marshal(path)
	pathString := string(pathJson)
	req := httplib.Get("http://" + conf.GetApiGateway() + "/viewWeatherForecast")
	req.Param("path", pathString)
	str, _ := req.Bytes()
	forecasts := model.WeatherForecasts{}
	_ = json.Unmarshal(str, &forecasts)
	fmt.Println(forecasts)
	size := len(forecasts.Rainy)
	forecastLists := []model.WeatherForecast{}
	for i := 0; i < size; i++ {
		forecast := model.WeatherForecast{
			Date:        forecasts.Date[i],
			Temperature: forecasts.Temperature[i],
			Cloudy:      forecasts.Cloudy[i],
			Rainy:       forecasts.Rainy[i],
			Wind:        forecasts.Wind[i],
			Humidity:    forecasts.Humidity[i],
		}
		forecastLists = append(forecastLists, forecast)
	}
	this.Data["forecasts"] = forecastLists
}
