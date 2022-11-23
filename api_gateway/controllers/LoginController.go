package controllers

import (
	"api_gateway/conf"
	"api_gateway/model"
	"api_gateway/proto"
	"context"
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"google.golang.org/grpc"
)

type LoginController struct {
	web.Controller
}

func Dial(host string) (*grpc.ClientConn, error) {
	fmt.Println("Dialing " + host)
	conn, err := grpc.Dial(host, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Println("")
	}
	return conn, nil
}
func (this *LoginController) CheckLogin(userId string) bool {
	fmt.Println("check login")
	if userId != "" {
		conn, _ := Dial(conf.GetConnectionConf("login"))
		client := proto.NewLoginServiceClient(conn)
		response, e := client.CheckLogin(context.TODO(), &proto.CheckRequest{UserId: userId})
		if e != nil {
			fmt.Println(e)
		}
		return response.IsLogged
	} else {
		return false
	}
}

func (this *LoginController) CheckUsername(userId string) bool {
	fmt.Println("check username")
	if userId != "" {
		conn, _ := Dial(conf.GetConnectionConf("login"))
		client := proto.NewLoginServiceClient(conn)
		response, e := client.CheckUsername(context.TODO(), &proto.CheckUsernameRequest{Username: userId})
		if e != nil {
			fmt.Println(e)
		}
		return response.IsRegistered
	} else {
		return false
	}
}

func (this *LoginController) login(credentialJson string) string {
	fmt.Println("LOGIN")
	conn, _ := Dial(conf.GetConnectionConf("login"))
	client := proto.NewLoginServiceClient(conn)
	fmt.Println("CONNECTION")
	var credential = model.Credential{}
	_ = json.Unmarshal([]byte(credentialJson), &credential)
	response, e := client.Login(context.TODO(), &proto.LoginRequest{Username: credential.Username, Password: credential.Password})
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println("SERVIZIO")
	fmt.Println(response.User)
	return response.User
}

func (this *LoginController) signIn(credentialJson string) string {
	fmt.Println("SIGNIN")
	conn, _ := Dial(conf.GetConnectionConf("login"))
	client := proto.NewLoginServiceClient(conn)
	fmt.Println("CONNECTION")
	var credential = model.Credential{}
	_ = json.Unmarshal([]byte(credentialJson), &credential)
	response, e := client.Signin(context.TODO(), &proto.LoginRequest{Username: credential.Username, Password: credential.Password})
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println("SERVIZIO")
	fmt.Println(response.User)
	return response.User
}

func (this *LoginController) LogOut(userId string) {
	fmt.Println("logout " + userId)
	conn, _ := Dial(conf.GetConnectionConf("login"))
	client := proto.NewLoginServiceClient(conn)
	_, e := client.LogOut(context.TODO(), &proto.CheckRequest{UserId: userId})
	if e != nil {
		fmt.Println(e)
	}

}

func (this *LoginController) Get() {

}

func (this *LoginController) Post() {
	fmt.Println("POST API GATEWAY")
	var credentialJson string
	var userId string
	if this.Ctx.Input.Query("signin") != "" {
		fmt.Println("SIGN IN")
		credentialJson = this.Ctx.Input.Query("signin")
		userId = this.signIn(credentialJson)
	} else if this.Ctx.Input.Query("login") != "" {
		fmt.Println("LOG IN")
		credentialJson = this.Ctx.Input.Query("login")
		userId = this.login(credentialJson)
	}
	fmt.Println(userId)
	this.Ctx.WriteString(userId)
}
