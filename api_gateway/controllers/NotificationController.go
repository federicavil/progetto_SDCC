package controllers

import (
	"api_gateway/grpc"
	"api_gateway/proto"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"strconv"
)

type NotificationController struct {
	web.Controller
}

func (this *NotificationController) Get() {
	userId := this.Ctx.Input.Query("userId")
	isLogged := this.Ctx.Input.Query("isLogged")
	if isLogged == "" {
		loginController := LoginController{}
		isLogged = strconv.FormatBool(loginController.CheckLogin(userId))
		this.Ctx.WriteString(isLogged)
	} else {
		//CIRCUIT BREAKER
		fmt.Println("GET NOTIFY")
		invites := grpc.GetInvites(proto.InviteInput{Username: userId})
		fmt.Println(string(invites))
		this.Ctx.WriteString(string(invites))
	}

}
