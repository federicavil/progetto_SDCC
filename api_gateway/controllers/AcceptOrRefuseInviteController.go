package controllers

import (
	"api_gateway/grpc"
	"api_gateway/proto"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"strconv"
)

type AcceptOrRefuseInviteController struct {
	web.Controller
}

func (this *AcceptOrRefuseInviteController) Get() {
	fmt.Println("InAddVisitControllerGet")
	userId := this.Ctx.Input.Query("userId")
	loginController := LoginController{}
	isLogged := loginController.CheckLogin(userId)
	this.Ctx.WriteString(strconv.FormatBool(isLogged))
}

func (this *AcceptOrRefuseInviteController) Post() {
	var visit proto.InviteResponse
	idVisit := this.Ctx.Input.Query("id_visit")
	username := this.Ctx.Input.Query("username")
	response := this.Ctx.Input.Query("response")
	visit.Username = &username
	visit.ID_Visit = &idVisit
	visit.Response = &response
	fmt.Println("InAddVisitControllerPost")

	ret := grpc.AcceptOrRefuseInvite(visit)
	defer this.ServeJSON()
	this.Ctx.WriteString(strconv.Itoa(ret))

}
