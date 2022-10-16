package main

import (
	_ "api_gateway/routers"
	"github.com/beego/beego/v2/server/web"
)

func main() {
	web.Run()
}
