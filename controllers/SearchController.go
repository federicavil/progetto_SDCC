package controllers

import "github.com/beego/beego/v2/server/web"

type SearchController struct {
	web.Controller
}

func (this *SearchController) Get() {
	this.TplName = "searchPath.html"
}
