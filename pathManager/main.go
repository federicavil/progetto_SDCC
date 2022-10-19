package main

import (
	"log"
	"net"
	"net/rpc"
	"pathManager/controller"
)

func main() {
	add := new(controller.Add)
	err := rpc.Register(add)
	if err != nil {
		return
	}
	//rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":8082")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	//go http.Serve(l, nil)
	rpc.Accept(l)
}
