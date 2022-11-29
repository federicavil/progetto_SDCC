package controllers

import (
	"frontend/conf"
	"frontend/model"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	"github.com/beego/beego/v2/client/httplib"
)

type AddNewVisitController struct {
	beego.Controller
	session session.Store
}

/*
* Prepare del client: verifica la presenza di nuove notifiche e imposta la view da mostrare all'utente
 */
func (this *AddNewVisitController) Prepare() {
	this.session = this.StartSession()
	notifications := this.session.Get("notifications")
	if notifications != nil {
		this.Data["newNotifications"] = true
	} else {
		this.Data["newNotifications"] = false
	}
	this.TplName = "addNewVisit.html"
}

/*
* Gestione chiamata GET: richiama CheckLogin per contattare API Gateway al fine di verificare che l'utente sia loggato
* Se non lo è, effettua redirect alla pagina di login.
* Mantiene anche informazioni sul percorso interessato alla visita da aggiungere
 */
func (this *AddNewVisitController) Get() {
	isLogged := CheckLogin(this.session, "/addNewVisit")
	if !isLogged {
		err := this.session.Set("prevPage", "addNewVisit")
		if err != nil {
			return
		}
		this.Redirect("login", 302)
	}
	this.Data["pathname"] = this.session.Get("selectedPath")
}

/*
* Gestione chiamata POST: Recupera informazioni da un form sulla view e invoca API Gateway al fine di aggiungere
* una nuova visita al sistema
 */
func (this *AddNewVisitController) Post() {
	//saveBtn := this.GetString("saveVisit")
	visitTime := this.GetString("visitTime")

	pathname := this.session.Get("selectedPath").(model.MountainPath)

	req := httplib.Post("http://" + conf.GetApiGateway() + "/addNewVisit")
	userid := this.session.Get("userId").(string)

	req.Param("pathname", pathname.Name)
	req.Param("timestamp", visitTime)
	req.Param("username", userid)
	_, err := req.Bytes()
	if err != nil {
		return
	}
	//this.session.Set("selectedPath", newPath)
	this.Data["pathname"] = this.session.Get("selectedPath")

}
