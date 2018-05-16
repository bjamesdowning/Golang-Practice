package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(go-mysql:3306)/test")
	if err != nil {
		log.Println(err)
	}
	return db
}
