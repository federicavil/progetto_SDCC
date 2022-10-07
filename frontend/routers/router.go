package routers

import (
	controllers "frontend/controllers"
	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/", &controllers.MainController{})
	web.Router("/searchPath", &controllers.SearchController{})
	web.Router("/assistedResearch", &controllers.AssistedResearchController{})
}
