package controllers

import (
	"fmt"
	"frontend/model"
	"github.com/astaxie/beego"
)

type ViewPathInfoController struct {
	beego.Controller
}

func (this *ViewPathInfoController) Prepare() {
	this.TplName = "viewPathInfo.html"
}

func (this *ViewPathInfoController) Get() {
	session := this.StartSession()
	path := session.Get("selectedPath").(model.MountainPath)
	this.Data["path"] = path
	fmt.Println(path)
}

func (this *ViewPathInfoController) Post() {
	if this.GetString("viewReviews") != "" {
		fmt.Println("btn view review")
		this.Redirect("viewReviews", 302)
	} else if this.GetString("addReview") != "" {
		fmt.Println("btn add review")
		this.Redirect("addReview", 302)
	} else if this.GetString("viewForecast") != "" {
		fmt.Println("btn view forecast")
		this.Redirect("viewWeatherForecast", 302)
	}

}
