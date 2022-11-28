package controllers

import (
	"api_gateway/grpc"
	"api_gateway/proto"
	"github.com/beego/beego/v2/server/web"
	"strconv"
)

type AddNewVisitController struct {
	web.Controller
}

/*
* Gestione chiamata GET: verifica che l'utente che esegue la chiamata sia loggato.
* @returns {string}: 1 se utente loggato, 0 se utente non loggato
 */
func (this *AddNewVisitController) Get() {
	userId := this.Ctx.Input.Query("userId")
	loginController := LoginController{}
	isLogged := loginController.CheckLogin(userId)
	this.Ctx.WriteString(strconv.FormatBool(isLogged))
}

/*
* Gestione chiamata POST: recupera parametri dalla chiamata REST e richiama la funzione AddNewVisit
* presente nella cartella grpc.
* @returns {string}: 1 se utente loggato, 0 se utente non loggato #TODO: RIVEDERE
 */
func (this *AddNewVisitController) Post() {
	var visit proto.InputVisit
	pathname := this.Ctx.Input.Query("pathname")
	username := this.Ctx.Input.Query("username")
	timestamp := this.Ctx.Input.Query("timestamp")
	visit.Username = &username
	visit.Pathname = &pathname
	visit.Timestamp = &timestamp

	ret := grpc.AddNewVisit(visit)
	defer this.ServeJSON()
	this.Ctx.WriteString(strconv.Itoa(ret))

}
