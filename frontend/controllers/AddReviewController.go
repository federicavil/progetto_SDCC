package controllers

import (
	"encoding/json"
	"fmt"
	"frontend/conf"
	"frontend/model"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	"github.com/beego/beego/v2/client/httplib"
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
		fmt.Println("back")
		this.Redirect("viewInfo", 302)
	} else {
		fmt.Println("POST REVIEW")
		var vote int
		/*for i := 5; i >= 1; i-- {
			star := string(i)
			if this.GetString(star) != "" {
				vote = i
				break
			}
		}*/
		btn1 := this.GetString("1")
		btn2 := this.GetString("2")
		btn3 := this.GetString("3")
		btn4 := this.GetString("4")
		btn5 := this.GetString("5")
		if btn5 != "" {
			vote = 5
		}
		if btn4 != "" {
			vote = 4
		}
		if btn3 != "" {
			vote = 3
		}
		if btn2 != "" {
			vote = 2
		}
		if btn1 != "" {
			vote = 1
		}

		fmt.Println(vote)
		path := this.session.Get("selectedPath").(model.MountainPath)

		review := model.Review{}
		review.Vote = vote
		review.Title = this.GetString("title")
		review.Comment = this.GetString("comment")
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
