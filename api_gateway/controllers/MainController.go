package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type MainController struct {
	web.Controller
}

func (this *MainController) Get() {
	//this.TplName = "home.html"

	//conn = httplib.GetRequest()
	//conn.request(method="POST", url=url, body=postdata, headers=headers)
	//resp = conn.getresponse()
}
