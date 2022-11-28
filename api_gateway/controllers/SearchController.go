package controllers

import (
	"api_gateway/conf"
	"api_gateway/model"
	"github.com/beego/beego/v2/server/web"

	"log"
	"net/http"
	"net/rpc"
)

type SearchController struct {
	web.Controller
}

/*
* Recupera informazioni su un percorso dato il suo nome
* @param {string}: nome del percorso
* @returns {[]byte}: lista di percorsi
 */
func SearchMountainPaths(name string) []byte {
	client, err := rpc.Dial("tcp", conf.GetConnectionConf("path_manager"))
	if err != nil {
		log.Fatal("dialing:", err)
	}
	param := model.SimpleSearchStruct{name}
	var results []byte
	//args := &search.Args{name}
	err = client.Call("Search.SimpleSearch", &param, &results)
	if err != nil {
		log.Fatal("SimpleSearch error:", err)
	}
	return results
}

/*
* Gestione chiamata POST: effettua ricerca dei sentieri usando come parametro quello ricevuto dalla chiamata REST
* @returns {string} lista di percorsi
 */
func (this *SearchController) Post() {
	defer this.ServeJSON()
	var paths []byte
	paths = SearchMountainPaths(this.Ctx.Input.Query("name"))
	this.Ctx.WriteString(string(paths))
	this.Ctx.Output.Status = http.StatusOK
}
