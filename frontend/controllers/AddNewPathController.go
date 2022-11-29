package controllers

import (
	"encoding/json"
	"frontend/conf"
	"frontend/model"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	"github.com/beego/beego/v2/client/httplib"
)

type AddNewPathController struct {
	beego.Controller
	session session.Store
}

/*
* Prepare del client: verifica la presenza di nuove notifiche e imposta la view da mostrare all'utente
 */
func (this *AddNewPathController) Prepare() {
	this.session = this.StartSession()
	notifications := this.session.Get("notifications")
	if notifications != nil {
		this.Data["newNotifications"] = true
	} else {
		this.Data["newNotifications"] = false
	}
	this.TplName = "addNewPath.html"
}

/*
* Gestione chiamata GET: richiama CheckLogin per contattare API Gateway al fine di verificare che l'utente sia loggato
* Se non lo è, effettua redirect alla pagina di login
 */
func (this *AddNewPathController) Get() {
	isLogged := CheckLogin(this.session, "/addNewPath")
	if !isLogged {
		err := this.session.Set("prevPage", "addPath")
		if err != nil {
			return
		}
		this.Redirect("login", 302)
	}
}

/*
* Gestione chiamata POST: Recupera informazioni da un form sulla view e invoca API Gateway al fine di aggiungere
* un nuovo percorso al sistema
 */
func (this *AddNewPathController) Post() {
	saveBtn := this.GetString("savePath")
	if saveBtn != "" {
		newPath := model.MountainPath{}
		err := this.ParseForm(&newPath)
		if err != nil {
			return
		}
		pathJson, _ := json.Marshal(newPath)
		pathString := string(pathJson)
		req := httplib.Post("http://" + conf.GetApiGateway() + "/addNewPath")
		req.Param("path", pathString)
		_, err = req.Bytes()
		if err != nil {
			return //TODO: popup?
		}
		this.session.Set("selectedPath", newPath)
		this.Redirect("viewInfo", 302)
	}
}
