package controllers

import (
	"encoding/json"
	"frontend/conf"
	"frontend/model"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	"github.com/beego/beego/v2/client/httplib"
)

type SearchController struct {
	beego.Controller
	pathlist []model.MountainPath
	session  session.Store
}

/*
* Prepare del client: verifica la presenza di nuove notifiche e imposta la view da mostrare all'utente
 */
func (this *SearchController) Prepare() {
	this.session = this.StartSession()
	notifications := this.session.Get("notifications")
	if notifications != nil {
		this.Data["newNotifications"] = true
	} else {
		this.Data["newNotifications"] = false
	}
	this.TplName = "searchPath.html"
}

/*
* Funzione che chiama API Gateway per effettuare ricerca per nome dei percorsi
* @param {string}: parametro di ricerca
* @returns {[]byte}: lista dei percorsi risultanti dalla ricerca in json
 */
func SimpleSearchPost(name string) []byte {
	req := httplib.Post("http://" + conf.GetApiGateway() + "/simplesearch")
	req.Param("name", name)
	str, _ := req.Bytes()
	//fmt.Println("STRINGA OTTENUTA DA APIGW: " + str)
	//strings.Index(str, name)
	return str
}

/*
* Gestione chiamata GET: recupera lista di percorsi dalla sessione e le mostra sulla view
 */
func (this *SearchController) Get() {
	paths := this.session.Get("pathList")
	if paths != nil {
		this.pathlist = paths.([]model.MountainPath)
		this.Data["paths"] = this.pathlist
	}
}

/*
* Gestione chiamata POST: Recupera informazioni da un form sulla view e invoca API Gateway al fine di effettuare
* una ricerca dei percorsi basate sul nome del percorso, oppure al fine di ottenere informazioni su un percorso
* specifico per poi mostrarle in una nuova view.
 */
func (this *SearchController) Post() {
	// Chiama il metodo di ricerca
	name := this.GetString("pathName")
	selected := this.GetString("path")
	viewAll := this.GetString("viewAll")
	if name != "" || viewAll != "" {
		pathlist := SimpleSearchPost(name)
		if pathlist != nil && string(pathlist[len(pathlist)-4:]) == "null" {
			pathlist = pathlist[:len(pathlist)-4]
		}
		var paths []model.MountainPath
		_ = json.Unmarshal([]byte(pathlist), &paths)
		this.Data["paths"] = paths
		err := this.session.Set("pathList", paths)
		if err != nil {
			return
		}
		this.TplName = "searchPath.html"
	}
	if selected != "" {
		selectedPath := model.MountainPath{}
		this.pathlist = this.session.Get("pathList").([]model.MountainPath)
		for i := 0; i < len(this.pathlist); i++ {
			if this.pathlist[i].Name == selected {
				selectedPath = this.pathlist[i]
				break
			}
		}
		err := this.session.Set("selectedPath", selectedPath)
		if err != nil {
			return
		}
		this.Redirect("viewInfo", 302)
	}
}
