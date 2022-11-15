package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net"
	"net/rpc"
	"pathManager/conf"
	"pathManager/controller"
	"strconv"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func pgTry() {
	// connection string
	config, err := conf.LoadConfig("./conf")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	port, _ := strconv.Atoi(config.Port)

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Host, port, config.User, config.Password, config.Dbname)
	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")

}

func main() {
	pgTry()
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
