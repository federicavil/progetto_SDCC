package routers

import (
	controllers "frontend/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/searchPath", &controllers.SearchController{})
	beego.Router("/assistedResearch", &controllers.AssistedResearchController{})
	beego.Router("/viewInfo", &controllers.ViewPathInfoController{})
	beego.Router("/addPath", &controllers.AddNewPathController{})
}
