package controllers

import (
	"fmt"
	"frontend/model"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
)

type ViewPathInfoController struct {
	beego.Controller
	session session.Store
}

func (this *ViewPathInfoController) Prepare() {
	this.session = this.StartSession()
	this.TplName = "viewPathInfo.html"
}

func (this *ViewPathInfoController) Get() {
	path := this.session.Get("selectedPath").(model.MountainPath)
	this.Data["path"] = path
	fmt.Println(path)
}

func (this *ViewPathInfoController) Post() {
	if this.GetString("viewReviews") != "" {
		fmt.Println("btn view review")
		this.Redirect("viewReviews", 302)
	} else if this.GetString("addNewVisit") != "" {
		fmt.Println("btn add visit")
		this.Redirect("addNewVisit", 302)
	} else if this.GetString("addReview") != "" {
		fmt.Println("btn add review")
		this.Redirect("addReview", 302)
	} else if this.GetString("viewForecast") != "" {
		fmt.Println("btn view forecast")
		this.Redirect("viewWeatherForecast", 302)
	}

}
