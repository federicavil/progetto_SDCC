package controllers

import (
	"frontend/model"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
)

type ViewPathInfoController struct {
	beego.Controller
	session session.Store
}

/*
* Prepare del client: verifica la presenza di nuove notifiche e imposta la view da mostrare all'utente
 */
func (this *ViewPathInfoController) Prepare() {
	this.session = this.StartSession()
	notifications := this.session.Get("notifications")
	if notifications != nil {
		this.Data["newNotifications"] = true
	} else {
		this.Data["newNotifications"] = false
	}
	this.TplName = "viewPathInfo.html"
}

/*
* Gestione chiamata GET: recupera informazioni su un percorso specifico e le mostra sulla view
 */
func (this *ViewPathInfoController) Get() {
	path := this.session.Get("selectedPath").(model.MountainPath)
	this.Data["path"] = path
}

/*
* Gestione chiamata POST: effettua redirect a una pagina specifica in base al pulsante premuto nella view:
*	- Lista recensioni di un certo percorso
*	- Aggiunta di una nuova visita nel percorso specifico
*	- Aggiunta di una nuova recensione del percorso specifico
*	- Visualizzazione delle previsioni del tempo dell'arco di una settimana nel percorso specifico
 */
func (this *ViewPathInfoController) Post() {
	if this.GetString("viewReviews") != "" {
		this.Redirect("viewReviews", 302)
	} else if this.GetString("addNewVisit") != "" {
		this.Redirect("addNewVisit", 302)
	} else if this.GetString("addReview") != "" {
		this.Redirect("addReview", 302)
	} else if this.GetString("viewForecast") != "" {
		this.Redirect("viewWeatherForecast", 302)
	}

}
