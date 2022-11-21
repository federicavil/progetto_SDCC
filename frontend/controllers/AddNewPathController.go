package controllers

import (
	"encoding/json"
	"fmt"
	"frontend/conf"
	"frontend/model"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	"github.com/beego/beego/v2/client/httplib"
)

type AddNewPathController struct {
	beego.Controller
	session session.Store
}

func (this *AddNewPathController) Prepare() {
	this.session = this.StartSession()
	notifications := this.session.Get("notifications")
	if notifications != nil {
		this.Data["newNotifications"] = true
	} else {
		this.Data["newNotifications"] = false
	}
	this.TplName = "addNewPath.html"
}

func (this *AddNewPathController) Get() {
	isLogged := CheckLogin(this.session, "/addNewPath")
	if !isLogged {
		err := this.session.Set("prevPage", "addPath")
		if err != nil {
			return
		}
		this.Redirect("login", 302)
	}
}

func (this *AddNewPathController) Post() {
	saveBtn := this.GetString("savePath")
	if saveBtn != "" {
		newPath := model.MountainPath{}
		err := this.ParseForm(&newPath)
		if err != nil {
			return
		}
		pathJson, _ := json.Marshal(newPath)
		pathString := string(pathJson)
		req := httplib.Post("http://" + conf.GetApiGateway() + "/addNewPath")
		req.Param("path", pathString)
		str, _ := req.Bytes()
		fmt.Println(str)
		this.session.Set("selectedPath", newPath)
		this.Redirect("viewInfo", 302)
	}
}
