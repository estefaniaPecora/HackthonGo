package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	StorageDB *sql.DB
)

func init() {
	dataSource := "root@tcp(localhost:3306)/dbproductos"

	var err error
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	StorageDB = db

	log.Println("DataBase Configured")

}
