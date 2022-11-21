package controllers

import (
	"api_gateway/grpc"
	"github.com/beego/beego/v2/server/web"
)

type GetAllVisitsController struct {
	web.Controller
}

func (this *GetAllVisitsController) Get() {
	username := this.Ctx.Input.Query("username")
	visits := grpc.GetAllVisits(username)
	defer this.ServeJSON()
	this.Data["json"] = string(visits)
}
