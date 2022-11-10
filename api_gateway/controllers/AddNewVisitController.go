package controllers

import (
	"api_gateway/grpc"
	"github.com/beego/beego/v2/server/web"
	"strconv"
)

type AddNewVisitController struct {
	web.Controller
}

func (this *AddNewVisitController) Post() {
	var visit grpc.InputVisit
	pathname := this.Ctx.Input.Query("pathname")
	username := this.Ctx.Input.Query("username")
	visit.Username = &username
	visit.Pathname = &pathname
	ret := grpc.AddNewVisit(visit)
	defer this.ServeJSON()
	this.Ctx.WriteString(strconv.Itoa(ret))

}
