package api_gw_funcs

import (
	"fmt"
	"frontend/controllers"
	"net/rpc"
)

func searchMountainPaths(name string) {
	client, _ := rpc.Dial("tcp", "localhost:8081")
	results := []controllers.MountainPath{}
	if err := client.Call("Handler.GetPaths", name, &results); err != nil {
		fmt.Printf("Error:1 %+v", err)
	} else {
		fmt.Printf("user found")
	}
}
