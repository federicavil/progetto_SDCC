package controllers

import (
	"api_gateway/grpc"
	"github.com/beego/beego/v2/server/web"
	"strconv"
)

type GetVisitByIDController struct {
	web.Controller
}

/*
  - Gestione chiamata GET: in base all'input, verifica che l'utente che esegue la chiamata sia loggato, oppure
  - richiama la funzione grpc.GetVisitById per ottenere una visita specifica
  - @returns {string}:
    -isLogged: 1 se utente loggato, 0 se utente non loggato
    -visit: json con informazioni sulla visita ottenuta
*/
func (this *GetVisitByIDController) Get() {
	userId := this.Ctx.Input.Query("userId")
	visitId := this.Ctx.Input.Query("visitId")
	if userId != "" {
		loginController := LoginController{}
		isLogged := loginController.CheckLogin(userId)
		this.Ctx.WriteString(strconv.FormatBool(isLogged))
	} else if visitId != "" {
		visit := grpc.GetVisitByID(visitId)
		this.Ctx.WriteString(string(visit))
	}
}
