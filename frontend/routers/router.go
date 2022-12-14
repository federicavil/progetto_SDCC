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
	beego.Router("/addNewVisit", &controllers.AddNewVisitController{})
	beego.Router("/viewVisits", &controllers.ViewVisitsController{})
	beego.Router("/viewVisitInfo", &controllers.ViewVisitInfoController{})
	beego.Router("/profile", &controllers.ProfileController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/addReview", &controllers.AddReviewController{})
	beego.Router("/viewReviews", &controllers.ViewReviewsController{})
	beego.Router("/viewWeatherForecast", &controllers.ViewWeatherForecastController{})
	beego.Router("/notifications", &controllers.NotificationController{})
}
