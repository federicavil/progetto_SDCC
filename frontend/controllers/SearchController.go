package controllers

import (
	"encoding/json"
	"frontend/model"
	"github.com/beego/beego/v2/client/httplib"
	"github.com/beego/beego/v2/server/web"
)

type SearchController struct {
	web.Controller
}

func (this *SearchController) Get() {
	this.TplName = "searchPath.html"
}

type Field struct {
	name  string
	value string
}

func SimplePost(name string) []byte {
	req := httplib.Post("http://127.0.0.1:5000/simplesearch")
	req.Param("name", name)
	data := []model.MountainPath{}
	str, _ := req.Bytes()
	_ = json.Unmarshal([]byte(str), &data)
	//fmt.Println("STRINGA OTTENUTA DA APIGW: " + str)
	//strings.Index(str, name)

	return str
}

func (this *SearchController) Post() {
	// Chiama il metodo di ricerca
	name := this.GetString("pathName")
	pathlist := SimplePost(name)
	pathlist = pathlist[:len(pathlist)-4]
	var paths []model.MountainPathRet
	_ = json.Unmarshal([]byte(pathlist), &paths)
	this.Data["paths"] = paths
	this.TplName = "searchPath.html"
}
