package controller

import (
	"database/sql"
	"log"
	"pathManager/conf"
	"pathManager/model"
	"strconv"
	"strings"
)

type Add int

/*
* Aggiunge un nuovo sentiero all'interno del database
* @Param {newPathPointer}: puntatore alla struttura contenente il il nuovo sentiero da aggiungere
* @Param {reply}: puntatore alla struttura in cui inserire la risposta
* @returns {error}
 */
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
	newPath.Name = strings.ReplaceAll(newPath.Name, "'", "''")
	newPath.Region = strings.ReplaceAll(newPath.Region, "'", "''")
	newPath.Province = strings.ReplaceAll(newPath.Province, "'", "''")
	newPath.City = strings.ReplaceAll(newPath.City, "'", "''")
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
