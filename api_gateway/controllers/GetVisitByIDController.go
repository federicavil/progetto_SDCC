package controllers

import (
	"api_gateway/grpc"
	"github.com/beego/beego/v2/server/web"
	"strconv"
)

type GetVisitByIDController struct {
	web.Controller
}

func (this *GetVisitByIDController) Get() {
	userId := this.Ctx.Input.Query("userId")
	visitId := this.Ctx.Input.Query("visitId")
	if userId != "" {
		loginController := LoginController{}
		isLogged := loginController.CheckLogin(userId)
		this.Ctx.WriteString(strconv.FormatBool(isLogged))
	} else if visitId != "" {
		visit := grpc.GetVisitByID(visitId)
		this.Ctx.WriteString(string(visit))
	}
}
