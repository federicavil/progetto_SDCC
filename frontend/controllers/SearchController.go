package controllers

import (
	"fmt"
	"frontend/api_gw_funcs"
	"github.com/beego/beego/v2/server/web"
)

type SearchController struct {
	web.Controller
}

func (this *SearchController) Get() {
	this.TplName = "searchPath.html"
}

func (this *SearchController) Post() {
	fmt.Println("POST CHIAMATO")
	// Chiama il metodo di ricerca
	name := this.GetString("pathName")
	pathlist := api_gw_funcs.SearchMountainPaths(name)
	this.Data["paths"] = pathlist
	this.TplName = "searchPath.html"
}
