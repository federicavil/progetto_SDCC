package main

import (
	_ "github.com/lib/pq"
	"log"
	"net"
	"net/rpc"
	"pathManager/conf"
	"pathManager/controller"
)

func main() {
	conf.DbConnect()
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
		log.Fatal("listen error: ", e)
	}
	rpc.Accept(l)
}
