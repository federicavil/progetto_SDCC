package controllers

import (
	"encoding/json"
	"fmt"
	"frontend/model"
	"github.com/astaxie/beego"
	"github.com/beego/beego/v2/client/httplib"
)

type AddNewPathController struct {
	beego.Controller
}

func (this *AddNewPathController) Prepare() {
	this.TplName = "addNewPath.html"

}

func (this *AddNewPathController) Get() {
	//TO DO : CONTROLLARE SE L'UTENTE E' LOGGATO E IN CASO MANDARLO SULLA PAGINA DI LOGIN

}

func (this *AddNewPathController) Post() {
	newPath := model.MountainPath{}
	err := this.ParseForm(&newPath)
	if err != nil {
		return
	}
	fmt.Println(newPath)

	newPathJson, _ := json.Marshal(newPath)
	req := httplib.Post("http://127.0.0.1:5000/addNewPath")
	req.Param("newPath", string(newPathJson))

	this.TplName = "addNewPath.html"
}
