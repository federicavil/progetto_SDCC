package controllers

import (
	"fmt"
	"github.com/beego/beego/v2/server/web"
)

type LoginController struct {
	web.Controller
}

func CheckLogin(userId string) bool {
	return userId == "1"
}

func login(credentialJson string) string {
	return "1"
}

func signIn(credentialJson string) string {
	return "1"
}

func (this *LoginController) Get() {

}

func (this *LoginController) Post() {
	var credentialJson string
	var userId string
	if this.Ctx.Input.Query("signin") != "" {
		fmt.Println("SIGN IN")
		credentialJson = this.Ctx.Input.Query("signin")
		userId = signIn(credentialJson)
	} else if this.Ctx.Input.Query("login") != "" {
		fmt.Println("LOG IN")
		credentialJson = this.Ctx.Input.Query("login")
		userId = login(credentialJson)
	}
	fmt.Println(userId)
	this.Ctx.WriteString(userId)
}
