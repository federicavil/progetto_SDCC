package routers

import (
	controllers "api_gateway/controllers"
	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/", &controllers.MainController{})
	web.Router("/simplesearch", &controllers.SearchController{})
	web.Router("/assistedsearch", &controllers.AssistedSearchController{})
	web.Router("/addNewPath", &controllers.AddNewPathController{})
	web.Router("/addNewVisit", &controllers.AddNewVisitController{})
	web.Router("/profile", &controllers.ProfileController{})
	web.Router("/login", &controllers.LoginController{})
}
