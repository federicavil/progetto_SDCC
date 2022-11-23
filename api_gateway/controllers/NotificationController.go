package controllers

import (
	"api_gateway/controllers/utils"
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
		cb := utils.GetCircuitBreaker()
		invites, err := cb.Execute(func() (interface{}, error) {
			invites, err := grpc.GetInvites(proto.InviteInput{Username: userId})
			return invites, err
		})
		if err != nil {
			fmt.Println("Notification error")
			invites, _ = json.Marshal("[]")
		}
		fmt.Println("GET NOTIFY")
		fmt.Println(string(invites.([]byte)))
		this.Ctx.WriteString(string(invites.([]byte)))
	}

}
