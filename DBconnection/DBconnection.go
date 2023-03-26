package DBconnection

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
)

var (
	DB              *sql.DB
	connectionError error
)

func Connect() {
	content, errcon := ioutil.ReadFile("D:\\desktop2\\GoProjects\\stdlib-golang-server(fixed)\\configs\\config.json")
	if errcon != nil {
		log.Fatal(errcon.Error())
	}

	var config map[string]string

	if err := json.Unmarshal(content, &config); err != nil {
		log.Fatalf("DBconnection (cant parse config): %v", err)
	}
	fmt.Print("conf: ")
	fmt.Println(config)

	var cfg = mysql.Config{
		User:                 config["User"],
		Passwd:               config["Password"],
		Net:                  config["NetworkConnection"],
		Addr:                 config["Address"],
		DBName:               config["Name"],
		AllowNativePasswords: true,
	}

	DB, connectionError = sql.Open("mysql", cfg.FormatDSN())
	if connectionError != nil {
		log.Fatalf("cannot connect to database (%v)", connectionError)
	}

	log.Println("Connected successfully")
}
