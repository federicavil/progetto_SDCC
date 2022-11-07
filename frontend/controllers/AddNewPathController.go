package controllers

import (
	"encoding/json"
	"fmt"
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
	req := httplib.Get("http://127.0.0.1:5000/addNewPath")
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
	newPath := model.MountainPath{}
	err := this.ParseForm(&newPath)
	if err != nil {
		return
	}
	pathJson, _ := json.Marshal(newPath)
	pathString := string(pathJson)
	req := httplib.Post("http://127.0.0.1:5000/addNewPath")
	req.Param("path", pathString)
	str, _ := req.Bytes()
	fmt.Println(str)
}
