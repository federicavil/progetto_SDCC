package controllers

import (
	"github.com/beego/beego/v2/server/web"
	"strconv"
)

type GetVisitByIDController struct {
	web.Controller
}

func (this *GetVisitByIDController) Get() {
	userId := this.Ctx.Input.Query("userId")
	loginController := LoginController{}
	isLogged := loginController.CheckLogin(userId)
	this.Ctx.WriteString(strconv.FormatBool(isLogged))
}
