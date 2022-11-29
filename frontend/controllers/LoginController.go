package controllers

import (
	"encoding/json"
	"fmt"
	"frontend/conf"
	"frontend/model"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	"github.com/beego/beego/v2/client/httplib"
	"strconv"
	"sync"
)

type LoginController struct {
	beego.Controller
	session session.Store
}

/*
* Prepare del client: imposta la view da mostrare all'utente
 */
func (this *LoginController) Prepare() {
	this.TplName = "login.html"
	this.session = this.StartSession()
}

func (this *LoginController) Get() {

}

/*
* Funzione che chiama API Gateway per effettuare il login/signin
* @param {string}: "login" o "signin"
* @param {model.Credential}: credenziali dell'utente
* @returns {string}: userId dell'utente autenticato/registrato
 */
func login(mode string, credential model.Credential) string {
	credentialJson, _ := json.Marshal(credential)
	req := httplib.Post("http://" + conf.GetApiGateway() + "/login")
	req.Param(mode, string(credentialJson))
	userId, _ := req.Bytes()
	return string(userId)
}

/*
* Funzione che chiama API Gateway per verificare che l'utente sia loggato
* @param {session.Store}: identificatore della sessione, usata per ricevere l'userId di sessione
* @param {string}: pagina relativa alla view che esegue la funzione
* @returns {bool}: true se check positivo, false se check negativo
 */
func CheckLogin(session session.Store, page string) bool {
	var userid string
	if session.Get("userId") == nil {
		userid = ""
	} else {
		userid = session.Get("userId").(string)
	}
	req := httplib.Get("http://" + conf.GetApiGateway() + page)
	req.Param("userId", userid)
	str, _ := req.Bytes()
	isLogged, _ := strconv.ParseBool(string(str))
	return isLogged
}

/*
* Gestione chiamata POST: Recupera informazioni da un form sulla view e invoca API Gateway al fine di effettuare
* l'operazione di login o di registrazione al sistema da parte di un utente.
* Dopo aver effettuato il login, viene eseguita una goroutine che riceve le notifiche dell'utente.
 */
func (this *LoginController) Post() {
	loginBtn := this.GetString("login")
	signinBtn := this.GetString("signin")
	credential := model.Credential{}
	err := this.ParseForm(&credential)
	loginForm := model.Login{}
	err = this.ParseForm(&loginForm)
	loginForm.Credential = credential
	if err != nil {
		return
	}
	userId := ""
	if signinBtn != "" {
		if loginForm.PasswordConfirmation != loginForm.Credential.Password {
			fmt.Println("ERRORE DA GESTIRE 1" + loginForm.PasswordConfirmation + " " + loginForm.Credential.Password)
			return
		} else {
			if len(loginForm.Credential.Password) <= 0 {
				fmt.Println("ERRORE DA GESTIRE 2")
				return
			}
		}
		userId = login("signin", loginForm.Credential)
	} else if loginBtn != "" {
		userId = login("login", loginForm.Credential)
	}
	fmt.Println(userId)
	if userId == "-1" {
		fmt.Println("login error")
		this.Data["login_error"] = "error"

	} else if userId != "" {
		error := this.session.Set("userId", userId)
		if error != nil {
			fmt.Println("Errore session")
			return
		}
		this.session.Set("sessionMutex", new(sync.Mutex))
		go NotificationPolling(userId, this.session)
		prevPage := this.session.Get("prevPage")
		this.Redirect(fmt.Sprint(prevPage), 302)
	}
}
