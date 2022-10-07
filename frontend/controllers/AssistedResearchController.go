package controllers

import (
	"frontend/api_gw_funcs"
	"frontend/model"
	"github.com/beego/beego/v2/server/web"
)

type AssistedResearchController struct {
	web.Controller
}

func (this *AssistedResearchController) Get() {
	this.TplName = "assistedResearch.html"
}

func (this *AssistedResearchController) Post() {
	filters := model.AdvancedSearchStruct{}
	err := this.ParseForm(&filters)
	if err != nil {
		return
	}
	pathList := api_gw_funcs.AdvancedSearchMountainPaths(filters)
	this.Data["pathList"] = pathList
	this.TplName = "assistedResearch.html"
}
