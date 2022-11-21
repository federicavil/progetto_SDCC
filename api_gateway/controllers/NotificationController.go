package controllers

import (
	"github.com/beego/beego/v2/server/web"
	"strconv"
)

type NotificationController struct {
	web.Controller
}

func (this *NotificationController) Get() {
	userId := this.Ctx.Input.Query("userId")
	loginController := LoginController{}
	isLogged := loginController.CheckLogin(userId)
	this.Ctx.WriteString(strconv.FormatBool(isLogged))
}
