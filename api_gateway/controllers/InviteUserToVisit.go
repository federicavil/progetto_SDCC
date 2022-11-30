package controllers

import (
	"api_gateway/grpc"
	"api_gateway/proto"
	"github.com/beego/beego/v2/server/web"
	"strconv"
)

type InviteUserToVisitController struct {
	loginController LoginController
	web.Controller
}

/*
* Gestione chiamata POST: recupera parametri dalla chiamata REST e, se l'utente da invitare esiste,
* esegue la funzione grpc.InviteUserToVisit per invitare una persona a una certa visita
* @returns {string}:
	* 0 se l'utente da invitare non esiste (errore)
	* -1 se #TODO: Vedere
	* -2 se l'utente è stato già invitato (errore)
	* -3 se l'invito è destinato all'autore stesso della visita (errore)
	* 1 se l'invito va a buon fine
*/
func (this *InviteUserToVisitController) Post() {
	var invite proto.Invite
	id_visit := this.Ctx.Input.Query("id_visit")
	username := this.Ctx.Input.Query("username")
	invite.Username = &username
	invite.ID_Visit = &id_visit
	isRegistrated := this.loginController.CheckUsername(username)
	ret := -10
	if isRegistrated {
		ret = grpc.InviteUserToVisit(invite)
	} else {
		ret = 0
	}
	defer this.ServeJSON()
	this.Ctx.WriteString(strconv.Itoa(ret))
}
