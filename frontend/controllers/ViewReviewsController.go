package controllers

import (
	"encoding/json"
	"frontend/conf"
	"frontend/model"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	"github.com/beego/beego/v2/client/httplib"
)

type ViewReviewsController struct {
	beego.Controller
	session session.Store
}

/*
* Prepare del client: verifica la presenza di nuove notifiche e imposta la view da mostrare all'utente
 */
func (this *ViewReviewsController) Prepare() {
	this.session = this.StartSession()
	notifications := this.session.Get("notifications")
	if notifications != nil {
		this.Data["newNotifications"] = true
	} else {
		this.Data["newNotifications"] = false
	}
	this.TplName = "viewReviews.html"
}

/*
* Gestione chiamata GET: comunica con API Gateway per ricevere la lista delle recensioni di un certo percorso
* dopodiché, mostra le recensioni nella view
 */
func (this *ViewReviewsController) Get() {
	path := this.session.Get("selectedPath").(model.MountainPath)
	pathName := path.Name
	req := httplib.Get("http://" + conf.GetApiGateway() + "/getReviews")
	req.Param("mountainPathName", pathName)
	str, _ := req.Bytes()

	var reviewsList []model.Review
	_ = json.Unmarshal(str, &reviewsList)
	this.Data["reviews"] = reviewsList
	this.Data["pathName"] = pathName

}

/*
* Gestione chiamata POST: Effettua un semplice redirect al dettaglio del percorso
 */
func (this *ViewReviewsController) Post() {
	this.Redirect("viewInfo", 302)
}
