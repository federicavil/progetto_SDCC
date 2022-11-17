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

func getUserProfile(userId string) string {
	fmt.Println("get user profile")
	conn, _ := Dial(conf.GetConnectionConf("login"))
	client := proto.NewLoginServiceClient(conn)
	response, e := client.GetProfile(context.TODO(), &proto.ProfileRequest{UserId: userId, Profile: ""})
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(response.Profile)
	return response.Profile
}

func setUserProfile(userId string, userProfile string) {
	fmt.Println("set user profile")
	conn, _ := Dial(conf.GetConnectionConf("login"))
	client := proto.NewLoginServiceClient(conn)
	_, e := client.UpdateProfile(context.TODO(), &proto.ProfileRequest{UserId: userId, Profile: userProfile})
	if e != nil {
		fmt.Println(e)
	}
}

func (this *ProfileController) Get() {
	userId := this.Ctx.Input.Query("userId")
	fmt.Println(userId)
	isLogged := this.loginController.CheckLogin(userId)
	var response string
	if isLogged {
		response = getUserProfile(userId)
	} else {
		response = "false"
	}
	this.Ctx.WriteString(response)
}

func (this *ProfileController) Post() {
	userId := this.Ctx.Input.Query("userId")
	userProfile := this.Ctx.Input.Query("userProfile")
	if userProfile == "" {
		this.loginController.LogOut(userId)
	} else {
		setUserProfile(userId, userProfile)
	}
}
