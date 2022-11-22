package controllers

import (
	"encoding/json"
	"frontend/conf"
	"frontend/model"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	"github.com/beego/beego/v2/client/httplib"
)

type SearchController struct {
	beego.Controller
	pathlist []model.MountainPath
	session  session.Store
}

func (this *SearchController) Prepare() {
	this.session = this.StartSession()
	notifications := this.session.Get("notifications")
	if notifications != nil {
		this.Data["newNotifications"] = true
	} else {
		this.Data["newNotifications"] = false
	}
	this.TplName = "searchPath.html"
}

type Field struct {
	name  string
	value string
}

func SimpleSearchPost(name string) []byte {
	req := httplib.Post("http://" + conf.GetApiGateway() + "/simplesearch")
	req.Param("name", name)
	str, _ := req.Bytes()
	//fmt.Println("STRINGA OTTENUTA DA APIGW: " + str)
	//strings.Index(str, name)
	return str
}

func (this *SearchController) Get() {
	paths := this.session.Get("pathList")
	if paths != nil {
		this.pathlist = paths.([]model.MountainPath)
		this.Data["paths"] = this.pathlist
	}
}

func (this *SearchController) Post() {
	// Chiama il metodo di ricerca
	name := this.GetString("pathName")
	selected := this.GetString("path")
	viewAll := this.GetString("viewAll")
	if name != "" || viewAll != "" {
		pathlist := SimpleSearchPost(name)
		if pathlist != nil && string(pathlist[len(pathlist)-4:]) == "null" {
			pathlist = pathlist[:len(pathlist)-4]
		}
		var paths []model.MountainPath
		_ = json.Unmarshal([]byte(pathlist), &paths)
		this.Data["paths"] = paths
		err := this.session.Set("pathList", paths)
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
