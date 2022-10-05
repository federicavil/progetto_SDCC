package api_gw_funcs

import (
	"fmt"
	"frontend/model"
	"log"
	"net/rpc"
)

type Args struct {
	Name string
}

func SearchMountainPaths(name string) []model.MountainPath {
	client, err := rpc.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	param := Args{name}
	results := []model.MountainPath{}
	//args := &search.Args{name}
	fmt.Println(param)
	err = client.Call("Search.SimpleSearch", &param, &results)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Println(results)
	fmt.Println(param)
	return results
}
