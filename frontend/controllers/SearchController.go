package controllers

import (
	"fmt"
	"frontend/api_gw_funcs"
	"frontend/model"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
)

type SearchController struct {
	beego.Controller
	pathlist []model.MountainPath
	session  session.Store
}

func (this *SearchController) Prepare() {
	this.session = this.StartSession()
	this.TplName = "searchPath.html"
}

func (this *SearchController) Get() {
	paths := this.session.Get("pathList")
	if paths != nil {
		this.pathlist = paths.([]model.MountainPath)
		this.Data["paths"] = this.pathlist
	}
}

func (this *SearchController) Post() {
	name := this.GetString("pathName")
	selected := this.GetString("path")
	fmt.Println("name " + name)
	fmt.Println("selected " + selected)
	if name != "" {
		this.pathlist = api_gw_funcs.SearchMountainPaths(name)
		this.Data["paths"] = this.pathlist
		err := this.session.Set("pathList", this.pathlist)
		if err != nil {
			return
		}
		this.TplName = "searchPath.html"
	}
	if selected != "" {
		selectedPath := model.MountainPath{}
		this.pathlist = this.session.Get("pathList").([]model.MountainPath)
		for i := 0; i < len(this.pathlist); i++ {
			if this.pathlist[i].Name == selected {
				selectedPath = this.pathlist[i]
				fmt.Println(selectedPath)
				break
			}
		}
		err := this.session.Set("selectedPath", selectedPath)
		if err != nil {
			return
		}
		this.Redirect("viewInfo", 302)
	}

}
