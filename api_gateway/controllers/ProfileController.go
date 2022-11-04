package controllers

import (
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"strconv"
)

type ProfileController struct {
	web.Controller
}

func getUserProfile(userId string) []byte {
	return []byte("prova")
}

func (this *ProfileController) Get() {
	userId := this.Ctx.Input.Query("userId")
	loginController := LoginController{}
	fmt.Println(userId)
	isLogged := loginController.CheckLogin(userId)
	if isLogged {
		//userProfile := getUserProfile(userId)
		//this.Ctx.WriteString(string(userProfile))
	}
	this.Ctx.WriteString(strconv.FormatBool(isLogged))
}

func (this *ProfileController) Post() {
	userId := this.Ctx.Input.Query("logOut")
	loginController := LoginController{}
	loginController.LogOut(userId)
}
