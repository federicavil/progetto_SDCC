package controllers

import (
	"encoding/json"
	"fmt"
	"frontend/conf"
	"frontend/model"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	"github.com/beego/beego/v2/client/httplib"
	"strconv"
)

type AddNewPathController struct {
	beego.Controller
	session session.Store
}

func (this *AddNewPathController) Prepare() {
	this.TplName = "addNewPath.html"
	this.session = this.StartSession()
}

func (this *AddNewPathController) Get() {
	var userid string
	if this.session.Get("userId") == nil {
		userid = ""
	} else {
		userid = this.session.Get("userId").(string)
	}
	req := httplib.Get("http://" + conf.GetApiGateway() + "/addNewPath")
	req.Param("userId", userid)
	str, _ := req.Bytes()
	isLogged, _ := strconv.ParseBool(string(str))
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
	addReviewBtn := this.GetString("addReview")
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
	} else if addReviewBtn != "" {
		this.Redirect("addReview", 302)
	}
}
