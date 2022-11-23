package controllers

import (
	"encoding/json"
	"fmt"
	"frontend/conf"
	"frontend/model"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/session"
	"strconv"
)

type ViewVisitsController struct {
	beego.Controller
	session session.Store
}

func GetAllVisits(id string) []byte {
	req := httplib.Post("http://" + conf.GetApiGateway() + "/getAllVisits")
	req.Param("id_visit", id)
	str, _ := req.Bytes()
	return str
}

func (this *ViewVisitsController) Prepare() {
	this.session = this.StartSession()
	this.TplName = "viewVisits.html"
}

func (this *ViewVisitsController) Get() {
	var userid string
	if this.session.Get("userId") == nil {
		userid = ""
	} else {
		userid = this.session.Get("userId").(string)
	}
	req := httplib.Get("http://" + conf.GetApiGateway() + "/getAllVisits")
	req.Param("userId", userid)
	str, _ := req.Bytes()
	isLogged, _ := strconv.ParseBool(string(str))
	if !isLogged {
		err := this.session.Set("prevPage", "viewVisits")
		if err != nil {
			return
		}
		this.Redirect("login", 302)
	} else {
		visits := GetAllVisits(userid)

		if visits != nil {
			if visits != nil && string(visits[len(visits)-4:]) == "null" {
				visits = visits[:len(visits)-4]
			}
			var vss []model.MountainVisitComplete
			fmt.Println(string(visits))
			json.Unmarshal(visits, &vss)
			fmt.Println(vss)

			err := this.session.Set("visitList", vss)
			if err != nil {
				return
			}

			this.Data["visits"] = vss
		}
	}

}

func (this *ViewVisitsController) Post() {
	visitId := this.GetString("visitId")
	print("POST DI VIEWVISITS con id: " + visitId)
	if visitId != "" {
		selectedVisit := model.MountainVisitComplete{}
		visitlist := this.session.Get("visitList").([]model.MountainVisitComplete)
		for i := 0; i < len(visitlist); i++ {
			if strconv.Itoa(visitlist[i].ID_Visit) == visitId {
				selectedVisit = visitlist[i]
				break
			}
		}
		err := this.session.Set("selectedVisit", selectedVisit)
		if err != nil {
			return
		}

		print("\nSONO IN VIEWVISITS E HO: ")
		fmt.Println(this.session.Get("selectedVisit"))
		this.Redirect("viewVisitInfo", 302)
	}
}
