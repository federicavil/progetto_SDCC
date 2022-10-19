package controllers

import (
	"api_gateway/model"
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"log"
	"net/rpc"
)

type AddNewPathController struct {
	web.Controller
}

func AddNewPath(newPath model.MountainPath) error {
	client, err := rpc.Dial("tcp", "127.0.0.1:8082")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var results []byte
	//args := &search.Args{name}
	err = client.Call("pathManager.AddNewPath", &newPath, &results)
	if err != nil {
		log.Fatal("SimpleSearch error:", err)
	}
	return err
}

func (this *AddNewPathController) Post() {
	defer this.ServeJSON()
	newPathJson := this.Ctx.Input.Query("newPath")
	newPath := model.MountainPath{}
	err := json.Unmarshal([]byte(newPathJson), &newPath)
	if err != nil {
		fmt.Println("Unmarshal error")
	}
	// Controllo se già c'è un sentiero con quel nome
	err = AddNewPath(newPath)
	if err != nil {
		fmt.Println("Add error")
	}

}
