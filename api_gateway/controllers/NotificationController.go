package controllers

import (
	"api_gateway/controllers/utils"
	"api_gateway/grpc"
	"api_gateway/proto"
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"strconv"
)

type NotificationController struct {
	web.Controller
}

/*
* Gestione chiamata GET: in base all'input, verifica che l'utente che esegue la chiamata sia loggato, oppure
* tramite circuit breaker chiama la funzione grpc.GetInvites per ottenere notifiche di inviti da parte di altri
* utenti
* @returns {string}:
    * isLogged: 1 se utente loggato, 0 se utente non loggato
    * invites: json con informazioni sugli inviti ricevuti
*/
func (this *NotificationController) Get() {
	userId := this.Ctx.Input.Query("userId")
	isLogged := this.Ctx.Input.Query("isLogged")
	if isLogged == "" {
		loginController := LoginController{}
		isLogged = strconv.FormatBool(loginController.CheckLogin(userId))
		this.Ctx.WriteString(isLogged)
	} else {
		//CIRCUIT BREAKER
		cb := utils.GetCircuitBreaker()
		invites, err := cb.Execute(func() (interface{}, error) {
			invites, err := grpc.GetInvites(proto.InviteInput{Username: userId})
			return invites, err
		})
		if err != nil {
			fmt.Println("Notification error")
			invites, _ = json.Marshal("[]")
		}
		this.Ctx.WriteString(string(invites.([]byte)))
	}

}
