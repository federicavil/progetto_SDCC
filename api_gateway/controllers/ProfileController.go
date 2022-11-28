package controllers

import (
	"api_gateway/conf"
	"api_gateway/proto"
	"context"
	"fmt"
	"github.com/beego/beego/v2/server/web"
)

type ProfileController struct {
	web.Controller
	loginController LoginController
}

func (this *ProfileController) Prepare() {
	this.loginController = LoginController{}
}

/*
* Recupera informazioni su un utente dato il suo userId
* @param {string}: userId dell'utente
* @returns {string}: info dell'utente in json
 */
func getUserProfile(userId string) string {
	conn, _ := Dial(conf.GetConnectionConf("login"))
	client := proto.NewLoginServiceClient(conn)
	response, e := client.GetProfile(context.TODO(), &proto.ProfileRequest{UserId: userId, Profile: ""})
	if e != nil {
		fmt.Println(e)
	}
	return response.Profile
}

/*
* Imposta informazioni di un utente dato il suo userId
* @param {string}: userId dell'utente
* @param {string}: info dell'utente in json
 */
func setUserProfile(userId string, userProfile string) {
	conn, _ := Dial(conf.GetConnectionConf("login"))
	client := proto.NewLoginServiceClient(conn)
	_, e := client.UpdateProfile(context.TODO(), &proto.ProfileRequest{UserId: userId, Profile: userProfile})
	if e != nil {
		fmt.Println(e)
	}
}

/*
* Gestione chiamata GET: verifica che l'utente che esegue la chiamata sia loggato, e se l'esito della verifica è
* positivo, recupera informazioni di profilo dell'utente.
* @returns {string}: se utente loggato, informazioni di profilo dell'utente in json, altrimenti false.
 */
func (this *ProfileController) Get() {
	userId := this.Ctx.Input.Query("userId")
	isLogged := this.loginController.CheckLogin(userId)
	var response string
	if isLogged {
		response = getUserProfile(userId)
	} else {
		response = "false"
	}
	this.Ctx.WriteString(response)
}

/*
* Gestione chiamata POST: in base ai parametri REST, effettua logout oppure aggiorna le informazioni del profilo utente
 */
func (this *ProfileController) Post() {
	defer this.ServeJSON()
	userId := this.Ctx.Input.Query("userId")
	userProfile := this.Ctx.Input.Query("userProfile")
	if userProfile == "" {
		this.loginController.LogOut(userId)
	} else {
		setUserProfile(userId, userProfile)
	}
}
