package controllers

import (
	"frontend/api_gw_funcs"
	"github.com/beego/beego/v2/server/web"
)

type SearchController struct {
	web.Controller
}

const (
	EE string = "EE"
)

func (this *SearchController) Get() {
	this.TplName = "searchPath.html"
}

func (this *SearchController) Post() {
	// Chiama il metodo di ricerca
	//path_1 := MountainPath{
	//	Name:     "Gran Sasso",
	//	Altitude: 1234,
	//	Location: Location{
	//		City:     "pippo",
	//		Province: "Aquila",
	//		Region:   "Abruzzo",
	//	},
	//	Length:     1094,
	//	Level:      "EE",
	//	Cyclable:   false,
	//	Family:     false,
	//	Historical: false,
	//}
	//path_2 := MountainPath{
	//	Name:     "Piccolo Sasso",
	//	Altitude: 1234,
	//	Location: Location{
	//		City:     "pippo",
	//		Province: "Aquila",
	//		Region:   "Abruzzo",
	//	},
	//	Length:     1094,
	//	Level:      "EE",
	//	Cyclable:   false,
	//	Family:     true,
	//	Historical: false,
	//}
	//pathList := []MountainPath{path_1, path_2}
	name := this.GetString("pathName")
	pathlist := api_gw_funcs.SearchMountainPaths(name)
	this.Data["paths"] = pathlist
	this.TplName = "searchPath.html"
}
