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

type ProfileController struct {
	beego.Controller
	session session.Store
}

func (this *ProfileController) Prepare() {
	this.session = this.StartSession()
	notifications := this.session.Get("notifications")
	if notifications != nil {
		this.Data["newNotifications"] = true
	} else {
		this.Data["newNotifications"] = false
	}
	this.TplName = "profile.html"
}

func (this *ProfileController) Get() {
	var userid string
	if this.session.Get("userId") == nil {
		userid = ""
	} else {
		userid = this.session.Get("userId").(string)
	}
	req := httplib.Get("http://" + conf.GetApiGateway() + "/profile")
	req.Param("userId", userid)
	str, _ := req.Bytes()
	fmt.Println("RESPONSE: " + string(str))
	if string(str) == "false" {
		err := this.session.Set("prevPage", "profile")
		if err != nil {
			return
		}
		this.Redirect("login", 302)
	} else if string(str) != "" {
		user := model.UserProfile{}
		_ = json.Unmarshal(str, &user)
		fmt.Println(user)
		this.Data["user"] = user
	}
}

func (this *ProfileController) Post() {
	logOutBtn := this.GetString("logOut")
	saveEditBtn := this.GetString("save")
	if logOutBtn != "" {
		req := httplib.Post("http://" + conf.GetApiGateway() + "/profile")
		userid := this.session.Get("userId").(string)
		req.Param("userId", userid)
		_, _ = req.Bytes()
		this.session.Set("userId", "")
		fmt.Println("Redirection")
		this.Redirect("login", 302)
	} else if saveEditBtn != "" {
		user := model.UserProfile{}
		err := this.ParseForm(&user)
		if err != nil {
			return
		}
		userJson, _ := json.Marshal(user)
		req := httplib.Post("http://" + conf.GetApiGateway() + "/profile")
		userid := this.session.Get("userId").(string)
		req.Param("userId", userid)
		req.Param("userProfile", string(userJson))
		_, _ = req.Bytes()
		this.Redirect("profile", 302)
	}

}
