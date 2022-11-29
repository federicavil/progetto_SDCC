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

/*
* Funzione che chiama API Gateway per effettuare ricerca avanzata dei percorsi
* @param {model.AssistedSearchStruct}: parametri di ricerca
* @returns {[]byte}: lista dei percorsi risultanti dalla ricerca in json
 */
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
	str, _ := req.Bytes()
	//fmt.Println("STRINGA OTTENUTA DA APIGW: " + str)
	//strings.Index(str, name)
	return str
}

/*
* Prepare del client: verifica la presenza di nuove notifiche e imposta la view da mostrare all'utente
 */
func (this *AssistedResearchController) Prepare() {
	this.session = this.StartSession()
	notifications := this.session.Get("notifications")
	if notifications != nil {
		this.Data["newNotifications"] = true
	} else {
		this.Data["newNotifications"] = false
	}
	this.TplName = "assistedResearch.html"
}

func (this *AssistedResearchController) Get() {

}

/*
* Gestione chiamata POST: Recupera informazioni da un form sulla view e invoca API Gateway al fine di effettuare
* una ricerca avanzata di percorsi nel sistema. Mostra i risultati sulla view
 */
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
