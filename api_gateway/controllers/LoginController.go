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

/*
* Funzione di utility che prepara la connessione grpc
* @param {string}: stringa contenente ip e porta da contattare
* @returns {*grpc.ClientConn}: connessione grpc
* @returns {error}
 */
func Dial(host string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(host, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Println("Dial connection error")
	}
	return conn, err
}

/*
* Verifica che un utente sia loggato
* @param {string}: userId dell'utente da verificare
* @returns {bool}: true se loggato, false se non loggato
 */
func (this *LoginController) CheckLogin(userId string) bool {
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

/*
* Verifica che un utente sia registrato
* @param {string}: userId dell'utente da verificare
* @returns {bool}: true se registrato, false se non registrato
 */
func (this *LoginController) CheckUsername(userId string) bool {
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

/*
* Effettua operazione di login dell'utente
* @param {string}: credenziali dell'utente che deve effettuare il login
* @returns {string}: userId dell'utente se il login è avvenuto con successo #TODO: verificare
 */
func (this *LoginController) login(credentialJson string) string {
	conn, _ := Dial(conf.GetConnectionConf("login"))
	client := proto.NewLoginServiceClient(conn)
	var credential = model.Credential{}
	_ = json.Unmarshal([]byte(credentialJson), &credential)
	response, e := client.Login(context.TODO(), &proto.LoginRequest{Username: credential.Username, Password: credential.Password})
	if e != nil {
		fmt.Println(e)
	}
	return response.User
}

/*
* Effettua operazione di registrazione dell'utente
* @param {string}: credenziali dell'utente che deve effettuare la registrazione
* @returns {string}: userId dell'utente se la registrazione è avvenuta con successo #TODO: verificare
 */
func (this *LoginController) signIn(credentialJson string) string {
	conn, _ := Dial(conf.GetConnectionConf("login"))
	client := proto.NewLoginServiceClient(conn)
	var credential = model.Credential{}
	_ = json.Unmarshal([]byte(credentialJson), &credential)
	response, e := client.Signin(context.TODO(), &proto.LoginRequest{Username: credential.Username, Password: credential.Password})
	if e != nil {
		fmt.Println(e)
	}
	return response.User
}

/*
* Effettua operazione di logout dell'utente
* @param {string}: userId dell'utente che deve effettuare il logout
 */
func (this *LoginController) LogOut(userId string) {
	conn, _ := Dial(conf.GetConnectionConf("login"))
	client := proto.NewLoginServiceClient(conn)
	_, e := client.LogOut(context.TODO(), &proto.CheckRequest{UserId: userId})
	if e != nil {
		fmt.Println(e)
	}

}

func (this *LoginController) Get() {
}

/*
* Gestione chiamata POST: recupera parametri dalla chiamata REST e, in base ai parametri, permette di fare login o
* registrazione alla piattaforma.
* @returns {string}: userID dell'utente loggato/registrato
 */
func (this *LoginController) Post() {
	var credentialJson string
	var userId string
	if this.Ctx.Input.Query("signin") != "" {
		credentialJson = this.Ctx.Input.Query("signin")
		userId = this.signIn(credentialJson)
	} else if this.Ctx.Input.Query("login") != "" {
		credentialJson = this.Ctx.Input.Query("login")
		userId = this.login(credentialJson)
	}
	this.Ctx.WriteString(userId)
}
