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

/*
* Funzione che chiama API Gateway per invitare un utente alla visita creata dall'utente loggato
* @param {string}: userId dell'utente da invitare
* @param {string}: id della visita
* @returns {[]byte}: risposta della chiamata, usata per stampare diversi messaggi su view in base all'esito dell'invito
 */
func InviteUserPost(userId string, id_visit string) []byte {
	req := httplib.Post("http://" + conf.GetApiGateway() + "/inviteUserToVisit")
	req.Param("username", userId)
	req.Param("id_visit", id_visit)
	str, _ := req.Bytes()
	return str
}

/*
* Prepare del client: imposta la view da mostrare all'utente
 */
func (this *ViewVisitInfoController) Prepare() {
	this.session = this.StartSession()
	this.TplName = "viewVisitInfo.html"
}

/*
* Gestione chiamata GET: richiama CheckLogin per contattare API Gateway al fine di verificare che l'utente sia loggato
* Se non lo è, effettua redirect alla pagina di login.
* Altrimenti, riceve informazioni su una visita specifica e la stampa su schermo
 */
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
		fmt.Println(this.Data["Phrase"])
		var phrase Phrase
		if this.session.Get("phrase") != nil {
			phrase = this.session.Get("phrase").(Phrase)
		} else {
			phrase = Phrase{Value: ""}
		}

		this.Data["Phrase"] = phrase
		fmt.Println(this.Data["Phrase"])
	}
}

type Phrase struct {
	Value string
}

/*
* Gestione chiamata POST: Recupera informazioni da un form sulla view e invoca API Gateway al fine di invitare
* un utente alla visita definita dall'organizzatore.
 */
func (this *ViewVisitInfoController) Post() {
	// Chiama il metodo di ricerca
	invitedUser := this.GetString("invitedUser")
	visit := this.session.Get("selectedVisit").(model.MountainVisitComplete)
	this.session.Set("phrase", "")
	this.session.Set("selectedVisit", visit)
	if invitedUser != "" {
		resp := InviteUserPost(invitedUser, strconv.Itoa(visit.ID_Visit))
		fmt.Println("\n\nGUARDA QUI1: " + string(resp) + "\n\n")

		if resp != nil && string(resp[len(resp)-4:]) == "null" {
			resp = resp[:len(resp)-4]
		}
		phrase := ""

		fmt.Println("\n\nGUARDA QUI2: " + string(resp) + "\n\n")
		if string(resp) == "-2" {
			phrase = "User already invited"
		} else if string(resp) == "-3" {
			phrase = "You can't invite yourself!"
		} else if string(resp) == "-4" {
			phrase = "SQS error. Retry later."
		} else if string(resp) == "0" {
			phrase = "User not found"
		} else if string(resp) == "1" {
			phrase = "User invited correctly"
		}
		this.session.Set("phrase", Phrase{phrase})
	}

	this.Redirect("viewVisitInfo", 302)

}
