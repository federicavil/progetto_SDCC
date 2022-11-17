package controllers

import (
	"encoding/json"
	"fmt"
	"frontend/conf"
	"frontend/model"
	"github.com/astaxie/beego"
	"github.com/beego/beego/v2/client/httplib"
)

type ViewReviewsController struct {
	beego.Controller
}

func (this *ViewReviewsController) Prepare() {
	this.TplName = "viewReviews.html"
}

func (this *ViewReviewsController) Get() {
	path := this.StartSession().Get("selectedPath").(model.MountainPath)
	pathName := path.Name
	fmt.Println("VIEW REVIEW " + pathName)
	req := httplib.Get("http://" + conf.GetApiGateway() + "/getReviews")
	req.Param("mountainPathName", pathName)
	str, _ := req.Bytes()

	var reviewsList []model.Review
	_ = json.Unmarshal(str, &reviewsList)
	this.Data["reviews"] = reviewsList
	this.Data["pathName"] = pathName

}

func (this *ViewReviewsController) Post() {
	this.Redirect("viewInfo", 302)
}
