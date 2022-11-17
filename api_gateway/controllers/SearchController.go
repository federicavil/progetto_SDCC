package controllers

import (
	"api_gateway/conf"
	"api_gateway/model"
	"fmt"
	"github.com/beego/beego/v2/server/web"

	"log"
	"net/http"
	"net/rpc"
)

type SearchController struct {
	web.Controller
}

func SearchMountainPaths(name string) []byte {
	client, err := rpc.Dial("tcp", conf.GetConnectionConf("path_manager"))
	if err != nil {
		log.Fatal("dialing:", err)
	}
	fmt.Println("DIAL")
	param := model.SimpleSearchStruct{name}
	var results []byte
	//args := &search.Args{name}
	err = client.Call("Search.SimpleSearch", &param, &results)
	if err != nil {
		log.Fatal("SimpleSearch error:", err)
	}
	return results
}

func (this *SearchController) Post() {
	fmt.Println("POST")
	defer this.ServeJSON()
	var paths []byte
	paths = SearchMountainPaths(this.Ctx.Input.Query("name"))
	this.Ctx.WriteString(string(paths))
	this.Ctx.Output.Status = http.StatusOK

	return
}
