package controllers

import (
	"frontend/api_gw_funcs"
	"frontend/model"
	"github.com/beego/beego/v2/server/web"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
)

type AssistedResearchController struct {
	beego.Controller
	session session.Store
}

func (this *AssistedResearchController) Prepare() {
	this.session = this.StartSession()
	this.TplName = "assistedResearch.html"
}

func (this *AssistedResearchController) Get() {

}

func (this *AssistedResearchController) Post() {
	filters := model.AdvancedSearchStruct{}
	err := this.ParseForm(&filters)
	if err != nil {
		return
	}
	pathList := api_gw_funcs.AdvancedSearchMountainPaths(filters)
	err = this.session.Set("pathList", pathList)
	if err != nil {
		return
	}
	this.Redirect("searchPath", 302)
}
