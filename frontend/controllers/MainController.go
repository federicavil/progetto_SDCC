package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

/*
* Prepare del client: imposta la view da mostrare all'utente
 */
func (this *MainController) Prepare() {
	this.TplName = "home.html"
}

func (this *MainController) Get() {

}
