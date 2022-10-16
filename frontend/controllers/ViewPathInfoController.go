package controllers

import (
	"fmt"
	"frontend/model"
	"github.com/astaxie/beego"
)

type ViewPathInfoController struct {
	beego.Controller
}

func (this *ViewPathInfoController) Prepare() {
	this.TplName = "viewPathInfo.html"
}

func (this *ViewPathInfoController) Get() {
	session := this.StartSession()
	path := session.Get("selectedPath").(model.MountainPath)
	this.Data["path"] = path
	fmt.Println(path)
}
