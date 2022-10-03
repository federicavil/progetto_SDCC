package main

import (
	"github.com/beego/beego/v2/server/web"
	_ "progetto_SDCC/routers"
)

func main() {
	web.Run()
}
