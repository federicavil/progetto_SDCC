package controllers

import (
	"api_gateway/model"
	"github.com/beego/beego/v2/server/web"
	"log"
	"net/http"
	"net/rpc"
)

type SearchController struct {
	web.Controller
}

func SearchMountainPaths(name string) []byte {
	client, err := rpc.Dial("tcp", "127.0.0.1:8081")
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

func AdvancedSearchMountainPaths(filters model.AdvancedSearchStruct) []model.MountainPath {
	client, err := rpc.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	results := []model.MountainPath{}
	//args := &search.Args{name}
	err = client.Call("Search.AdvancedSearch", &filters, &results)
	if err != nil {
		log.Fatal("AdvancedSearch error:", err)
	}
	return results
}

func (this *SearchController) Post() {
	defer this.ServeJSON()
	//this.TplName = "home.html"
	//var stringhetta = "AOAOAO"
	var paths []byte
	paths = SearchMountainPaths(this.Ctx.Input.Query("name"))
	//jsn, err := json.Marshal(&paths)
	//if err != nil {
	//	log.Fatal("Unmarshal error:", err)
	//}
	//fmt.Println(jsn)
	this.Ctx.WriteString(string(paths))
	this.Ctx.Output.Status = http.StatusOK

	return
}
