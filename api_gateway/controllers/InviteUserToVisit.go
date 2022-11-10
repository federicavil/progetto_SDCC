package controllers

import (
	"api_gateway/grpc"
	"github.com/beego/beego/v2/server/web"
	"strconv"
)

type InviteUserToVisitController struct {
	web.Controller
}

func (this *InviteUserToVisitController) Post() {
	var invite grpc.Invite
	id_visit := this.Ctx.Input.Query("id_visit")
	username := this.Ctx.Input.Query("username")
	invite.Username = &username
	invite.ID_Visit = &id_visit
	ret := grpc.InviteUserToVisit(invite)
	defer this.ServeJSON()
	this.Ctx.WriteString(strconv.Itoa(ret))
}
