package main

import (
	_ "frontend/routers"
	"github.com/beego/beego/v2/server/web"
)

func main() {
	web.Run()
}
