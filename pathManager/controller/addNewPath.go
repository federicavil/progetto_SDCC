package controller

import (
	"database/sql"
	"fmt"
	"log"
	"pathManager/conf"
	"pathManager/model"
	"strconv"
)

type Add int

func (t *Add) AddNewPath(newPathPointer *model.MountainPath, reply *[]byte) error {
	var db, err, quote = conf.DbConnect()
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db) // Defer Closing the database

	newPath := *newPathPointer
	query := `INSERT INTO ` + quote + `Path` + quote + ` VALUES ('` +
		newPath.Name + "', '" +
		strconv.Itoa(newPath.Altitude) + "', '" +
		strconv.Itoa(newPath.Length) + "', '" +
		newPath.Level + "', '" +
		strconv.Itoa(t.boolToInt(newPath.Cyclable)) + "', '" +
		strconv.Itoa(t.boolToInt(newPath.Family)) + "', '" +
		strconv.Itoa(t.boolToInt(newPath.Historical)) + "', '" +
		newPath.Region + "', '" +
		newPath.Province + "', '" +
		newPath.City + "')"

	_, err = db.Exec(query)
	if err != nil {
		fmt.Println("Errore query: ")
		log.Fatal(err)
	}
	*reply = []byte("Okay")
	return err
}

func (t *Add) boolToInt(myVar bool) int {
	if myVar {
		return 1
	}
	return 0
}
