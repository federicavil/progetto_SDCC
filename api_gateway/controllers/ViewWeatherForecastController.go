package controllers

import (
	"api_gateway/grpc"
	"api_gateway/proto"
	"fmt"
	"github.com/beego/beego/v2/server/web"
)

type ViewWeatherForecastController struct {
	web.Controller
}

func (this *ViewWeatherForecastController) Get() {
	mountainpath := this.Ctx.Input.Query("path")
	fmt.Println("GET FORECAST " + mountainpath)

	result := grpc.GetForecast(proto.ForecastInput{Path: mountainpath})
	fmt.Println(result)
	//response := proto.ForecastOutput{
	//	Forecasts: "{\"time\" : [\"11/11/2022\",\"12/11/2022\",\"13/11/2022\"],\"temperature\" : [4,7,12],\"cloud_cover\" : [40,100,75],\"precipitation\" : [9,10,0],\"wind_speed\" : [12,3,5],\"humidity\" : [4,8,2]}",
	//}
	/*conn, _ := Dial("9092")
	client := proto.NewWeatherForecastServiceClient(conn)
	response, e := client.GetForecast(context.TODO(), &proto.ForecastInput{Path: mountainpath})
	if e != nil {
		fmt.Println(e)
	}*/
	fmt.Println(result.Forecasts)
	this.Ctx.WriteString(result.Forecasts)
}
