package controllers

import (
	"encoding/json"
	"fmt"
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
	this.TplName = "profile.html"
	this.session = this.StartSession()
}

func (this *ProfileController) Get() {
	var userid string
	if this.session.Get("userId") == nil {
		userid = ""
	} else {
		userid = this.session.Get("userId").(string)
	}
	req := httplib.Get("http://127.0.0.1:5000/profile")
	req.Param("userId", userid)
	str, _ := req.Bytes()
	fmt.Println("RESPONSE: " + string(str))
	if string(str) == "false" {
		err := this.session.Set("prevPage", "profile")
		if err != nil {
			return
		}
		this.Redirect("login", 302)
	} else {
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
		req := httplib.Post("http://127.0.0.1:5000/profile")
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
		req := httplib.Post("http://127.0.0.1:5000/profile")
		userid := this.session.Get("userId").(string)
		req.Param("userId", userid)
		req.Param("userProfile", string(userJson))
		_, _ = req.Bytes()
	}

}
