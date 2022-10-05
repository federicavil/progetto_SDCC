package main

import (
	"log"
	"net"
	"net/rpc"
	"search/searchPath"
)

func main() {
	search := new(searchPath.Search)
	err := rpc.Register(search)
	if err != nil {
		return
	}
	//rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":8081")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	//go http.Serve(l, nil)
	rpc.Accept(l)
}
