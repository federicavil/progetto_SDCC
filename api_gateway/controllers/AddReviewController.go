package controllers

import (
	"api_gateway/conf"
	"api_gateway/proto"
	"context"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"strconv"
)

type AddReviewController struct {
	web.Controller
}

/*
* Gestione chiamata GET: verifica che l'utente che esegue la chiamata sia loggato.
* @returns {string}: 1 se utente loggato, 0 se utente non loggato
 */
func (this *AddReviewController) Get() {
	userId := this.Ctx.Input.Query("userId")
	loginController := LoginController{}
	isLogged := loginController.CheckLogin(userId)
	this.Ctx.WriteString(strconv.FormatBool(isLogged))
}

/*
* Gestione chiamata POST: recupera parametri dalla chiamata REST ed effettua la chiamata gRPC al microservizio
* ReviewManager, in particolare richiama la funzione AddReview
* @returns {string}: 1 se utente loggato, 0 se utente non loggato # TODO:RIVEDERE
 */
func (this *AddReviewController) Post() {
	review := this.Ctx.Input.Query("review")
	conn, _ := Dial(conf.GetConnectionConf("review_manager"))
	client := proto.NewReviewManagerServiceClient(conn)
	response, err := client.AddReview(context.TODO(), &proto.AddReviewRequest{Review: review})
	if err != nil {
		fmt.Println(err)
	}
	this.Ctx.WriteString(response.AddedReview)
}
