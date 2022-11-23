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

func InviteUserPost(username string, id_visit string) []byte {
	req := httplib.Post("http://" + conf.GetApiGateway() + "/inviteUserToVisit")
	req.Param("username", username)
	req.Param("id_visit", id_visit)
	str, _ := req.Bytes()
	//fmt.Println("STRINGA OTTENUTA DA APIGW: " + str)
	//strings.Index(str, name)
	return str
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
		this.Data["visit"] = visit
		if visit.Creator == this.session.Get("userId") {
			this.Data["isCreator"] = true
		} else {
			this.Data["isCreator"] = false
		}
		this.Data["phrase"] = ""
	}
}

type Phrase struct {
	Value string
}

func (this *ViewVisitInfoController) Post() {
	// Chiama il metodo di ricerca
	invitedUser := this.GetString("invitedUser")
	visit := this.session.Get("selectedVisit").(model.MountainVisitComplete)

	this.session.Set("selectedVisit", visit)
	if invitedUser != "" {
		resp := InviteUserPost(invitedUser, strconv.Itoa(visit.ID_Visit))
		if resp != nil && string(resp[len(resp)-4:]) == "null" {
			resp = resp[:len(resp)-4]
		}
		phrase := ""

		if string(resp) == "-2" {
			phrase = "User already invited"
		} else if string(resp) == "-3" {
			phrase = "You can't invite yourself!"
		} else if string(resp) == "0" {
			phrase = "User not found"
		} else if string(resp) == "1" {
			phrase = "User invited correctly"
		}
		fmt.Println(this.Data["Phrase"])
		this.Data["Phrase"] = Phrase{Value: phrase}
		fmt.Println(this.Data["Phrase"])

	}

	this.Redirect("viewVisitInfo", 302)

}
