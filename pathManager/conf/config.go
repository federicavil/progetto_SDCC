package conf

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"gopkg.in/ini.v1"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Host     string `mapstructure:"host"`
	Dbname   string `mapstructure:"dbname"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Port     string `mapstructure:"port"`
	Quote    string `mapstructure:"quote"`
	Dbtype   string `mapstructure:"dbtype"`
}

func LoadIni(path string) (config Config, err error) {
	cfg, err := ini.Load(path)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	section := cfg.Section(cfg.Section("").Key("app_mode").String())
	config = Config{
		section.Key("host").String(),
		section.Key("dbname").String(),
		section.Key("user").String(),
		section.Key("password").String(),
		section.Key("port").String(),
		section.Key("quote").String(),
		section.Key("dbtype").String(),
	}
	return
}

func DbConnect() (*sql.DB, error, string) {
	// connection string
	config, err := LoadIni("conf/database.ini")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	port, _ := strconv.Atoi(config.Port)
	var conn string
	if config.Dbtype == "postgres" {
		conn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Host, port, config.User, config.Password, config.Dbname)
	} else {
		conn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.User, config.Password, config.Host, port, config.Dbname)

	}
	db, err := sql.Open(config.Dbtype, conn)
	CheckError(err)

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")
	return db, err, config.Quote
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
