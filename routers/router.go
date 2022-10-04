package routers

import (
	"github.com/beego/beego/v2/server/web"
	controllers2 "progetto_SDCC/controllers"
)

func init() {
	web.Router("/", &controllers2.MainController{})
	web.Router("/searchPath", &controllers2.SearchController{})
}
