package controllers

import (
	"encoding/json"
	"fmt"
	"frontend/model"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	"github.com/beego/beego/v2/client/httplib"
)

type LoginController struct {
	beego.Controller
	session session.Store
}

func (this *LoginController) Prepare() {
	this.TplName = "login.html"
	this.session = this.StartSession()
}

func (this *LoginController) Get() {

}

func login(mode string, credential model.Credential) string {
	credentialJson, _ := json.Marshal(credential)
	req := httplib.Post("http://127.0.0.1:5000/login")
	req.Param(mode, string(credentialJson))
	response, _ := req.Bytes()
	fmt.Println(string(response))
	return string(response)
}

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
		prevPage := this.session.Get("prevPage")
		this.Redirect(fmt.Sprint(prevPage), 302)
	}

}
