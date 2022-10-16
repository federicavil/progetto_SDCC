package api_gw_funcs

import (
	"fmt"
	"frontend/model"
	"log"
	"net/rpc"
)

func SearchMountainPaths(name string) []model.MountainPath {
	client, err := rpc.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	param := model.SimpleSearchStruct{name}
	results := []model.MountainPath{}
	//args := &search.Args{name}
	err = client.Call("Search.SimpleSearch", &param, &results)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	return results
}

func AdvancedSearchMountainPaths(filters model.AdvancedSearchStruct) []model.MountainPath {
	client, err := rpc.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	results := []model.MountainPath{}
	//args := &search.Args{name}
	err = client.Call("Search.AdvancedSearch", &filters, &results)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Println(results)
	return results
}
