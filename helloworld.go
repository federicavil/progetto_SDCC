package main

import "github.com/beego/beego/v2/server/web"

type MainController struct {
	web.Controller
}

func (this *MainController) Get() {
	this.TplName = "home.html"
}

func main() {
	web.Router("/", &MainController{})
	web.Run()
}
