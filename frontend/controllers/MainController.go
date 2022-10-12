package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Prepare() {
	this.TplName = "home.html"
}

func (this *MainController) Get() {

}
