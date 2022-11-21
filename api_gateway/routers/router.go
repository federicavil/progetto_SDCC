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
	web.Router("/acceptOrRefuseInvite", &controllers.AcceptOrRefuseInviteController{})
	web.Router("/inviteUserToVisit", &controllers.InviteUserToVisitController{})
	web.Router("/getAllVisits", &controllers.GetAllVisitsController{})
	web.Router("/getAllInvites", &controllers.GetAllInvitesController{})
	web.Router("/profile", &controllers.ProfileController{})
	web.Router("/login", &controllers.LoginController{})
	web.Router("/addReview", &controllers.AddReviewController{})
	web.Router("/getReviews", &controllers.ViewReviewsController{})
	web.Router("/viewWeatherForecast", &controllers.ViewWeatherForecastController{})
	web.Router("/notifications", &controllers.NotificationController{})
}
