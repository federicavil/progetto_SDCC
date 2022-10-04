package routers

import (
	controllers2 "frontend/controllers"
	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/", &controllers2.MainController{})
	web.Router("/searchPath", &controllers2.SearchController{})
}
