package controller

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
	"log"
)

type MountainPath struct {
	Name       string
	Altitude   int
	City       string
	Province   string
	Region     string
	Length     int
	Level      string
	Cyclable   bool
	Family     bool
	Historical bool
}

type Args struct {
	Name string
}

type Search int

func (t *Search) SimpleSearch(args *Args, reply *[]MountainPath) error {
	//path_1 := MountainPath{
	//	Name:     "Gran Sasso",
	//	Altitude: 1234,
	//	Location: Location{
	//		City:     "pippo",
	//		Province: "Aquila",
	//		Region:   "Abruzzo",
	//	},
	//	Length:     1094,
	//	Level:      "EE",
	//	Cyclable:   false,
	//	Family:     false,
	//	Historical: false,
	//}
	//path_2 := MountainPath{
	//	Name:     "Piccolo Sasso",
	//	Altitude: 1234,
	//	Location: Location{
	//		City:     "pippo",
	//		Province: "Aquila",
	//		Region:   "Abruzzo",
	//	},
	//	Length:     1094,
	//	Level:      "EE",
	//	Cyclable:   false,
	//	Family:     true,
	//	Historical: false,
	//}
	var db, _ = sql.Open("sqlite3", "./search.db") // Open the created SQLite File
	defer db.Close()                               // Defer Closing the database

	var query = "SELECT * FROM Paths WHERE name LIKE " + "'%" + args.Name + "%'"
	fmt.Println(query)
	row, err := db.Query(query)
	if err != nil {
		fmt.Println("Errore query: ")
		log.Fatal(err)
	}
	defer row.Close()
	var path = MountainPath{}
	for row.Next() { // Iterate and fetch the records from result cursor
		row.Scan(&path.Name, &path.Altitude, &path.Length, &path.Level,
			&path.Cyclable, &path.Family, &path.Historical,
			&path.Region, &path.Province, &path.City)
		*reply = append(*reply, path)
	}
	fmt.Println(*reply)
	return nil
}
