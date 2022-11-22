package controllers

import (
	"api_gateway/grpc"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"net/http"
	"strconv"
)

type GetAllVisitsController struct {
	web.Controller
}

func (this *GetAllVisitsController) Get() {
	userId := this.Ctx.Input.Query("userId")
	loginController := LoginController{}
	isLogged := loginController.CheckLogin(userId)
	this.Ctx.WriteString(strconv.FormatBool(isLogged))
}

func (this *GetAllVisitsController) Post() {
	id := this.Ctx.Input.Query("id_visit")
	visits := grpc.GetAllVisits(id)
	fmt.Println(string(visits))
	defer this.ServeJSON()
	this.Ctx.WriteString(string(visits))
	this.Ctx.Output.Status = http.StatusOK
	//this.Data["json"] = string(visits)
}
