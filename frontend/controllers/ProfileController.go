package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	"github.com/beego/beego/v2/client/httplib"
	"strconv"
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
	isLogged, _ := strconv.ParseBool(string(str))
	if !isLogged {
		err := this.session.Set("prevPage", "profile")
		if err != nil {
			return
		}
		this.Redirect("login", 302)
	}
}

func (this *ProfileController) Post() {
	logOutBtn := this.GetString("logOut")
	if logOutBtn != "" {
		req := httplib.Post("http://127.0.0.1:5000/profile")
		userid := this.session.Get("userId").(string)
		req.Param("logOut", userid)
		_, _ = req.Bytes()
		this.session.Set("userId", "")
		fmt.Println("Redirection")
		this.Redirect("login", 302)
	}

}
