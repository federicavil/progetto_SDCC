package controllers

import (
	"api_gateway/grpc"
	"api_gateway/proto"
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/server/web"
)

type GetAllInvitesController struct {
	web.Controller
}

func (this *GetAllInvitesController) Get() {
	username := this.Ctx.Input.Query("username")
	invites := grpc.GetInvites(proto.InviteInput{Username: username})
	fmt.Println(string(invites))

	defer this.ServeJSON()
	var reply []byte
	reply, _ = json.Marshal(string(invites))
	this.Data["json"] = reply
	//this.Ctx.WriteString(string(invites))

}
