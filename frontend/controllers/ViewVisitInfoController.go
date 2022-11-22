package controllers

import (
	"fmt"
	"frontend/conf"
	"frontend/model"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/session"
	"strconv"
)

type ViewVisitInfoController struct {
	beego.Controller
	session session.Store
}

func (this *ViewVisitInfoController) Prepare() {
	this.session = this.StartSession()
	this.TplName = "viewVisitInfo.html"
}

func (this *ViewVisitInfoController) Get() {
	var userid string
	if this.session.Get("userId") == nil {
		userid = ""
	} else {
		userid = this.session.Get("userId").(string)
	}
	req := httplib.Get("http://" + conf.GetApiGateway() + "/getVisitByID")
	req.Param("userId", userid)
	str, _ := req.Bytes()
	isLogged, _ := strconv.ParseBool(string(str))
	if !isLogged {
		err := this.session.Set("prevPage", "viewVisitInfo")
		if err != nil {
			return
		}
		this.Redirect("login", 302)
	} else {
		visit := this.session.Get("selectedVisit").(model.MountainVisitComplete)
		fmt.Println(visit)
		this.Data["visit"] = visit
	}

}
