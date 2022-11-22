package controllers

import (
	"fmt"
	"frontend/model"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
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
	isLogged := CheckLogin(this.session, "/getVisitByID")
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
