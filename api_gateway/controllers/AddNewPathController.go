package controllers

import (
	"api_gateway/conf"
	"api_gateway/model"
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"log"
	"net/rpc"
	"strconv"
)

type AddNewPathController struct {
	web.Controller
}

/*
* Effettua la chiamata goRPC al microservizio PathManager per aggiungere un nuovo path al sistema.
* @param {model.MountainPath}: Informazioni relative al path da salvare
* @returns {error}
 */
func AddNewPath(newPath model.MountainPath) error {
	client, err := rpc.Dial("tcp", conf.GetConnectionConf("path_manager"))
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var results []byte
	//args := &search.Args{name}
	err = client.Call("Add.AddNewPath", &newPath, &results)
	if err != nil {
		log.Fatal("AddPath error:", err)
	}
	return err
}

/*
* Gestione chiamata GET: verifica che l'utente che esegue la chiamata sia loggato.
* @returns {string}: 1 se utente loggato, 0 se utente non loggato
 */
func (this *AddNewPathController) Get() {
	userId := this.Ctx.Input.Query("userId")
	loginController := LoginController{}
	isLogged := loginController.CheckLogin(userId)
	this.Ctx.WriteString(strconv.FormatBool(isLogged))
}

/*
* Gestione chiamata POST: recupera parametro "path" dalla chiamata REST e richiama la funzione AddNewPath.
 */
func (this *AddNewPathController) Post() {
	defer this.ServeJSON()
	newPathJson := this.Ctx.Input.Query("path")
	newPath := model.MountainPath{}
	err := json.Unmarshal([]byte(newPathJson), &newPath)
	if err != nil {
		fmt.Println("Unmarshal error")
	}
	// Controllo se già c'è un sentiero con quel nome TODO: OCCHIO
	err = AddNewPath(newPath)
	if err != nil {
		fmt.Println("Add error")
	}
}
