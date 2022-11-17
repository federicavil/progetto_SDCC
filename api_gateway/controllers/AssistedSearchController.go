package controllers

import (
	"api_gateway/conf"
	"api_gateway/model"
	"github.com/beego/beego/v2/server/web"
	"log"
	"net/http"
	"net/rpc"
	"strconv"
)

type AssistedSearchController struct {
	web.Controller
}

func AssistedSearchMountainPaths(filters model.AdvancedSearchStruct) []byte {
	client, err := rpc.Dial("tcp", conf.GetConnectionConf("path_manager"))
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var results []byte
	//args := &search.Args{name}
	err = client.Call("Search.AdvancedSearch", &filters, &results)
	if err != nil {
		log.Fatal("AdvancedSearch error:", err)
	}
	return results
}

func (this *AssistedSearchController) Post() {
	defer this.ServeJSON()
	var paths []byte
	var filters model.AdvancedSearchStruct
	filters.City = this.Ctx.Input.Query("city")
	filters.Province = this.Ctx.Input.Query("province")
	filters.Region = this.Ctx.Input.Query("region")
	filters.Level = this.Ctx.Input.Query("level")
	filters.Cyclable, _ = strconv.Atoi(this.Ctx.Input.Query("cyclable"))
	filters.Family, _ = strconv.Atoi(this.Ctx.Input.Query("familySuitable"))
	filters.Historical, _ = strconv.Atoi(this.Ctx.Input.Query("historicalElements"))
	paths = AssistedSearchMountainPaths(filters)
	this.Ctx.WriteString(string(paths))
	this.Ctx.Output.Status = http.StatusOK
	return
}
