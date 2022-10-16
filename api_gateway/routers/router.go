package routers

import (
	controllers "api_gateway/controllers"
	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/", &controllers.MainController{})
	web.Router("/simplesearch", &controllers.SearchController{})
}
