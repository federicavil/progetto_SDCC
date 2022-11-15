package controllers

import (
	"api_gateway/grpc"
	"api_gateway/proto"
	"github.com/beego/beego/v2/server/web"
	"strconv"
)

type AddNewVisitController struct {
	web.Controller
}

func (this *AddNewVisitController) Get() {
	userId := this.Ctx.Input.Query("userId")
	loginController := LoginController{}
	isLogged := loginController.CheckLogin(userId)
	this.Ctx.WriteString(strconv.FormatBool(isLogged))
}

func (this *AddNewVisitController) Post() {
	var visit proto.InputVisit
	pathname := this.Ctx.Input.Query("pathname")
	username := this.Ctx.Input.Query("username")
	timestamp := this.Ctx.Input.Query("timestamp")
	visit.Username = &username
	visit.Pathname = &pathname
	visit.Timestamp = &timestamp
	ret := grpc.AddNewVisit(visit)
	defer this.ServeJSON()
	this.Ctx.WriteString(strconv.Itoa(ret))

}
