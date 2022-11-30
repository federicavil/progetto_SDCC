package main

import (
	_ "github.com/lib/pq"
	"net"
	"net/rpc"
	"pathManager/conf"
	"pathManager/controller"
	"time"
)

func main() {
	time.Sleep(3 * time.Second)
	i := 1
	for i < 10 {
		_, err, _ := conf.DbConnect()
		if err != nil {
			i++
			time.Sleep(time.Second)
		} else {
			i = 10
		}
	}
	add := new(controller.Add)
	err := rpc.Register(add)
	if err != nil {
		return
	}
	search := new(controller.Search)
	err = rpc.Register(search)
	if err != nil {
		return
	}
	l, e := net.Listen("tcp", ":8081")
	if e != nil {
		print("listen error: ", e)
	}
	rpc.Accept(l)
}
