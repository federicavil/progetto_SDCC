package controllers

import (
	"api_gateway/grpc"
	"api_gateway/proto"
	"github.com/beego/beego/v2/server/web"
	"strconv"
)

type AcceptOrRefuseInviteController struct {
	web.Controller
}

/*
- Gestione chiamata GET: verifica che l'utente che esegue la chiamata sia loggato.
* @returns {string}: 1 se utente loggato, 0 se utente non loggato
*/
func (this *AcceptOrRefuseInviteController) Get() {
	userId := this.Ctx.Input.Query("userId")
	loginController := LoginController{}
	isLogged := loginController.CheckLogin(userId)
	this.Ctx.WriteString(strconv.FormatBool(isLogged))
}

/*
- Gestione chiamata POST: recupera parametri dalla chiamata REST e richiama la funzione che effettuerà
- la chiamata gRPC al microservizio VisitManager per accettare/rifiutare l'invito.
* @returns {string}: 1 se utente loggato, 0 se utente non loggato #TODO: rivedere
*/
func (this *AcceptOrRefuseInviteController) Post() {
	var visit proto.InviteResponse
	idVisit := this.Ctx.Input.Query("id_visit")
	username := this.Ctx.Input.Query("username")
	response := this.Ctx.Input.Query("response")
	visit.Username = &username
	visit.ID_Visit = &idVisit
	visit.Response = &response

	ret := grpc.AcceptOrRefuseInvite(visit)
	defer this.ServeJSON()
	this.Ctx.WriteString(strconv.Itoa(ret))

}
