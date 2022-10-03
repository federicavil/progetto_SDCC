package routers

import (
	"github.com/beego/beego/v2/server/web"
	"progetto_SDCC/controllers"
)

func init() {
	web.Router("/", &controllers.MainController{})
}
