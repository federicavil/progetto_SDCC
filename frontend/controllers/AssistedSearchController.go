package controllers

import (
	"encoding/json"
	"fmt"
	"frontend/conf"
	"frontend/model"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	"github.com/beego/beego/v2/client/httplib"
	"strconv"
)

type AssistedResearchController struct {
	beego.Controller
	session session.Store
}

//City       string ` form:"city"`
//Province   string ` form:"province"`
//Region     string ` form:"region"`
//Level      string ` form:"level"`
//Cyclable   int    ` form:"cyclable"`
//Family     int    ` form:"familySuitable"`
//Historical int    ` form:"historicalElements"`

func AssistedPost(filters model.AssistedSearchStruct) []byte {
	req := httplib.Post("http://" + conf.GetApiGateway() + "/assistedsearch")
	fmt.Println(filters)
	req.Param("city", filters.City)
	req.Param("province", filters.Province)
	req.Param("region", filters.Region)
	req.Param("level", filters.Level)
	req.Param("cyclable", strconv.Itoa(filters.Cyclable))
	req.Param("familySuitable", strconv.Itoa(filters.Family))
	req.Param("historicalElements", strconv.Itoa(filters.Historical))
	data := []model.MountainPath{}
	str, _ := req.Bytes()
	_ = json.Unmarshal([]byte(str), &data)
	//fmt.Println("STRINGA OTTENUTA DA APIGW: " + str)
	//strings.Index(str, name)
	return str
}

func (this *AssistedResearchController) Prepare() {
	this.session = this.StartSession()
	this.TplName = "assistedResearch.html"
}

func (this *AssistedResearchController) Get() {

}

func (this *AssistedResearchController) Post() {
	filters := model.AssistedSearchStruct{}
	err := this.ParseForm(&filters)
	if err != nil {
		return
	}
	pathlist := AssistedPost(filters)
	fmt.Println(pathlist)
	if pathlist != nil {
		pathlist = pathlist[:len(pathlist)-4]
	}
	var paths []model.MountainPath
	_ = json.Unmarshal([]byte(pathlist), &paths)
	this.Data["paths"] = paths
	err = this.session.Set("pathList", paths)
	if err != nil {
		return
	}
	this.Redirect("searchPath", 302)
}
