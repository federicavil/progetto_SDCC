package controllers

import (
	"fmt"
	"frontend/model"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	"github.com/beego/beego/v2/client/httplib"
	"strconv"
)

type AddNewVisitController struct {
	beego.Controller
	session session.Store
}

func (this *AddNewVisitController) Prepare() {
	this.TplName = "addNewVisit.html"
	this.session = this.StartSession()
}

func (this *AddNewVisitController) Get() {
	var userid string
	if this.session.Get("userId") == nil {
		userid = ""
	} else {
		userid = this.session.Get("userId").(string)
	}
	req := httplib.Get("http://127.0.0.1:5000/addNewVisit")
	req.Param("userId", userid)
	str, _ := req.Bytes()
	isLogged, _ := strconv.ParseBool(string(str))
	if !isLogged {
		err := this.session.Set("prevPage", "addNewVisit")
		if err != nil {
			return
		}
		this.Redirect("login", 302)
	}
	this.Data["pathname"] = this.session.Get("selectedPath")

}

func (this *AddNewVisitController) Post() {
	//saveBtn := this.GetString("saveVisit")
	visitTime := this.GetString("visitTime")

	newPath := model.MountainVisit{}
	err := this.ParseForm(&newPath)

	if err != nil {
		return
	}
	pathname := this.session.Get("selectedPath").(model.MountainPath)
	fmt.Println(pathname.Name)

	//pathJson, _ := json.Marshal(newPath)
	//pathString := string(pathJson)
	req := httplib.Post("http://127.0.0.1:5000/addNewVisit")
	userid := this.session.Get("userId").(string)
	fmt.Println("VISITTIME: " + visitTime)
	//fmt.Println("NAME: " + pathname)
	newPath.Username = userid
	newPath.Pathname = pathname.Name
	req.Param("pathname", pathname.Name)
	req.Param("timestamp", visitTime)
	req.Param("username", userid)
	str, _ := req.Bytes()
	fmt.Println(str)
	//this.session.Set("selectedPath", newPath)
	this.Data["pathname"] = this.session.Get("selectedPath")

}
