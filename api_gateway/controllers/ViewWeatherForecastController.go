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

/*
* Gestione chiamata GET: effettua chiamata grpc per recuperare le previsioni del tempo settimanali di un dato percorso
* @returns {string}: lista delle previsioni, giorno per giorno, per una settimana in json
 */
func (this *ViewWeatherForecastController) Get() {
	mountainpath := this.Ctx.Input.Query("path")

	conn, _ := Dial(conf.GetConnectionConf("weather_forecast"))
	client := proto.NewWeatherForecastServiceClient(conn)
	response, e := client.GetForecast(context.TODO(), &proto.ForecastInput{Path: mountainpath})
	if e != nil {
		fmt.Println(e)
	}
	this.Ctx.WriteString(response.Forecasts)
}
