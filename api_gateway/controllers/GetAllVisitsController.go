package controllers

import (
	"api_gateway/grpc"
	"github.com/beego/beego/v2/server/web"
	"net/http"
	"strconv"
)

type GetAllVisitsController struct {
	web.Controller
}

/*
* Gestione chiamata GET: verifica che l'utente che esegue la chiamata sia loggato
* @returns {string}: 1 se utente loggato, 0 se utente non loggato
 */
func (this *GetAllVisitsController) Get() {
	userId := this.Ctx.Input.Query("userId")
	loginController := LoginController{}
	isLogged := loginController.CheckLogin(userId)
	this.Ctx.WriteString(strconv.FormatBool(isLogged))
}

/*
* Gestione chiamata POST: recupera parametri dalla chiamata REST ed esegue la funzione GetAllVisits
* presente nella cartella grpc
* @returns {string}: lista di tutte le visite a cui un utente partecipa
 */
func (this *GetAllVisitsController) Post() {
	id := this.Ctx.Input.Query("id_visit")
	visits := grpc.GetAllVisits(id)
	defer this.ServeJSON()
	this.Ctx.WriteString(string(visits))
	this.Ctx.Output.Status = http.StatusOK
	//this.Data["json"] = string(visits)
}
