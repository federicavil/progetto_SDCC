package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"log"
	"pathManager/conf"
	"pathManager/model"
	"strconv"
)

type Args struct {
	Name string
}

type Search int

/*
* Esegue una ricerca per nome di un sentiero
* @Param {args}: puntatore alla struttura contenente il nome del sentieor da cercare
* @Param {reply}: puntatore alla struttura in cui inserire la risposta
* @returns {error}
 */
func (t *Search) SimpleSearch(args *Args, reply *[]byte) error {
	var db, _, quote = conf.DbConnect()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db) // Defer Closing the database
	var query string
	if args.Name == "" {
		query = `SELECT * FROM ` + quote + `Path` + quote
	} else {
		query = `SELECT * FROM ` + quote + `Path` + quote + ` WHERE UPPER(name) LIKE UPPER(` + `'%` + args.Name + `%')`
	}
	fmt.Println(query)
	row, err := db.Query(query)
	if err != nil {
		fmt.Println("Errore query: ")
		log.Fatal(err)
	}
	defer func(row *sql.Rows) {
		err := row.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(row)
	var path = model.MountainPath{}
	var paths = []model.MountainPath{}
	for row.Next() { // Iterate and fetch the records from result cursor
		err := row.Scan(&path.Name, &path.Altitude, &path.Length, &path.Level,
			&path.Cyclable, &path.Family, &path.Historical,
			&path.Region, &path.Province, &path.City)
		if err != nil {
			return err
		}
		paths = append(paths, path)
	}
	*reply, _ = json.Marshal(paths)
	return nil
}

/*
* Esegue una ricerca per filtri di un sentiero
* @Param {pathreq}: puntatore alla struttura contenente la definizione dei filtri da applicare nella ricerca
* @Param {reply}: puntatore alla struttura in cui inserire la risposta
* @returns {error}
 */
func (t *Search) AdvancedSearch(pathreq *model.AdvancedSearchStruct, reply *[]byte) error {
	var db, _, quote = conf.DbConnect()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db) // Defer Closing the database

	if pathreq.Family == -1 {
		pathreq.Family = 0
	} else if pathreq.Family == 0 {
		pathreq.Family = -1
	}
	if pathreq.Cyclable == -1 {
		pathreq.Cyclable = 0
	} else if pathreq.Cyclable == 0 {
		pathreq.Cyclable = -1
	}
	if pathreq.Historical == -1 {
		pathreq.Historical = 0
	} else if pathreq.Historical == 0 {
		pathreq.Historical = -1
	}

	var query = `SELECT * FROM ` + quote + `Path` + quote
	if pathreq.City != "" || pathreq.Province != "" || pathreq.Region != "" ||
		pathreq.Level != "" || pathreq.Cyclable != -1 || pathreq.Family != -1 || pathreq.Historical != -1 {
		query = query + " WHERE"
	}
	var and int

	if pathreq.City != "" {
		if and == 0 {
			and = 1
		} else {
			query = query + " AND"
		}
		query = query + " UPPER(city) LIKE UPPER('%" + pathreq.City + "%')"
	}
	if pathreq.Province != "" {
		if and == 0 {
			and = 1
		} else {
			query = query + " AND"
		}
		query = query + " UPPER(province) LIKE UPPER('%" + pathreq.Province + "%')"
	}
	if pathreq.Region != "" {
		if and == 0 {
			and = 1
		} else {
			query = query + " AND"
		}
		query = query + " UPPER(region) LIKE UPPER('%" + pathreq.Region + "%')"
	}

	if pathreq.Level != "" {
		if and == 0 {
			and = 1
		} else {
			query = query + " AND"
		}
		query = query + " UPPER(level) LIKE UPPER('%" + pathreq.Level + "%')"
	}
	if pathreq.Cyclable != -1 {
		if and == 0 {
			and = 1
		} else {
			query = query + " AND"
		}
		query = query + " cyclable = '" + strconv.Itoa(pathreq.Cyclable) + "'"
	}
	if pathreq.Family != -1 {
		if and == 0 {
			and = 1
		} else {
			query = query + " AND"
		}
		query = query + " family = '" + strconv.Itoa(pathreq.Family) + "'"
	}
	if pathreq.Historical != -1 {
		if and == 0 {
			and = 1
		} else {
			query = query + " AND"
		}
		query = query + " historical = '" + strconv.Itoa(pathreq.Historical) + "'"
	}

	row, err := db.Query(query)
	if err != nil {
		fmt.Println("Errore query: ")
		log.Fatal(err)
	}
	defer func(row *sql.Rows) {
		err := row.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(row)
	var path = model.MountainPath{}
	var paths = []model.MountainPath{}
	for row.Next() { // Iterate and fetch the records from result cursor
		err := row.Scan(&path.Name, &path.Altitude, &path.Length, &path.Level,
			&path.Cyclable, &path.Family, &path.Historical,
			&path.Region, &path.Province, &path.City)
		if err != nil {
			return err
		}
		paths = append(paths, path)
	}
	*reply, _ = json.Marshal(paths)
	return err
}
