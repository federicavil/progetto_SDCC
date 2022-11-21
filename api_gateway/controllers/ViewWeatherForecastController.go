package controllers

import (
	"api_gateway/conf"
	"api_gateway/proto"
	"context"
	"fmt"
	"github.com/beego/beego/v2/server/web"
)

type ViewWeatherForecastController struct {
	web.Controller
}

func (this *ViewWeatherForecastController) Get() {
	mountainpath := this.Ctx.Input.Query("path")
	fmt.Println("GET FORECAST " + mountainpath)

	conn, _ := Dial(conf.GetConnectionConf("weather_forecast"))
	client := proto.NewWeatherForecastServiceClient(conn)
	response, e := client.GetForecast(context.TODO(), &proto.ForecastInput{Path: mountainpath})
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println("OBTAINED FORECAST: ")
	fmt.Println(response.Forecasts)
	this.Ctx.WriteString(response.Forecasts)
}
