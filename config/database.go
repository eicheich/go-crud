package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnnectDB() {

	db, err := sql.Open("mysql", "root:@/b_crudgo?parseTime=True")
	if err != nil {
		panic(err)
	}

	log.Println("Connected to database")
	DB = db
}
