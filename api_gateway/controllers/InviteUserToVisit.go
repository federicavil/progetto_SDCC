package controllers

import (
	"api_gateway/grpc"
	"api_gateway/proto"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"strconv"
)

type InviteUserToVisitController struct {
	loginController LoginController
	web.Controller
}

func (this *InviteUserToVisitController) Post() {
	var invite proto.Invite
	fmt.Println("\nINVITING...\n")
	id_visit := this.Ctx.Input.Query("id_visit")
	username := this.Ctx.Input.Query("username")
	invite.Username = &username
	invite.ID_Visit = &id_visit
	isLogged := this.loginController.CheckLogin(username)
	ret := -1
	if isLogged {
		ret = grpc.InviteUserToVisit(invite)
	} else {
		ret = 0
	}
	defer this.ServeJSON()
	this.Ctx.WriteString(strconv.Itoa(ret))
}
