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

/*
* Funzione che chiama API Gateway per ricevere tutte le visite relative a un utente
* @param {string}: userId dell'utente da invitare
* @returns {[]byte}: lista delle visite in json
 */
func GetAllVisits(id string) []byte {
	req := httplib.Post("http://" + conf.GetApiGateway() + "/getAllVisits")
	req.Param("id_visit", id)
	str, _ := req.Bytes()
	return str
}

/*
* Prepare del client: imposta la view da mostrare all'utente
 */
func (this *ViewVisitsController) Prepare() {
	this.session = this.StartSession()
	this.TplName = "viewVisits.html"
}

/*
* Gestione chiamata GET: richiama CheckLogin per contattare API Gateway al fine di verificare che l'utente sia loggato
* Se non lo è, effettua redirect alla pagina di login.
* Altrimenti, contatta API Gateway per ottenere tutte le visite relative a un userId e le mostra nella view
 */
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

/*
* Gestione chiamata POST: Recupera informazioni da un form sulla view e invoca API Gateway al fine di ricevere
* il dettaglio di una certa visita. Dopodiché viene fatto redirect alla view che mostra il dettaglio
 */
func (this *ViewVisitsController) Post() {
	visitId := this.GetString("visitId")
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

		this.Redirect("viewVisitInfo", 302)
	}
}
