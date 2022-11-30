package controllers

import (
	"encoding/json"
	"fmt"
	"frontend/conf"
	"frontend/model"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/session"
)

type AddReviewController struct {
	beego.Controller
	session session.Store
}

/*
* Prepare del client: verifica la presenza di nuove notifiche e imposta la view da mostrare all'utente
 */
func (this *AddReviewController) Prepare() {
	this.session = this.StartSession()
	notifications := this.session.Get("notifications")
	if notifications != nil {
		this.Data["newNotifications"] = true
	} else {
		this.Data["newNotifications"] = false
	}
	this.TplName = "addReview.html"
}

/*
* Gestione chiamata GET: richiama CheckLogin per contattare API Gateway al fine di verificare che l'utente sia loggato
* Se non lo è, effettua redirect alla pagina di login.
 */
func (this *AddReviewController) Get() {
	isLogged := CheckLogin(this.session, "/addReview")
	if !isLogged {
		err := this.session.Set("prevPage", "addReview")
		if err != nil {
			return
		}
		this.Redirect("login", 302)
	}
}

/*
* Gestione chiamata POST: Recupera informazioni da un form sulla view e invoca API Gateway al fine di aggiungere
* una nuova recensione relativa a un percorso al sistema
 */
func (this *AddReviewController) Post() {
	if this.GetString("backInfo") != "" {
		this.Redirect("viewInfo", 302)
	} else {
		var formResult model.ReviewForm
		err := this.ParseForm(&formResult)
		if err != nil {
			return
		}

		path := this.session.Get("selectedPath").(model.MountainPath)

		review := model.Review{}
		review.Vote = formResult.Vote
		review.Title = formResult.Title
		review.Comment = formResult.Comment
		review.Author = this.session.Get("userId").(string)
		review.MountainPathName = path.Name

		reviewJson, _ := json.Marshal(review)
		fmt.Println(string(reviewJson))
		req := httplib.Post("http://" + conf.GetApiGateway() + "/addReview")
		req.Param("review", string(reviewJson))
		str, _ := req.Bytes()
		fmt.Println(str)
	}

}
