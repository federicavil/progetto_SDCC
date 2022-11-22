package controllers

import (
	"api_gateway/grpc"
	"api_gateway/proto"
	"encoding/json"
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
		invites := grpc.GetInvites(proto.InviteInput{Username: userId})
		fmt.Println(string(invites))
		defer this.ServeJSON()
		var reply []byte
		reply, _ = json.Marshal(string(invites))
		this.Ctx.WriteString(string(reply))
	}

}
