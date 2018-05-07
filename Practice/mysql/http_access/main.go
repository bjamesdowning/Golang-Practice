package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

//using docker container running mysql 5.7
//test db by created from origination, hello table created from another program
//docker command:
//docker run -p 3306:3306 --name go-mysql -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=test -d mysql:5.7
func main() {
	db, err = sql.Open("mysql", "root:password@tcp(:3306)/test")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/CMDB", readAll)
	err := http.ListenAndServe(":8080", nil)
	check(err)
}

func readAll(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`SELECT * FROM test.CMDB;`)
	check(err)

	// data to be used in query
	var s, k, v string
	s = "RETRIEVED RECORDS:\n"

	// query
	for rows.Next() {
		err = rows.Scan(&k, &v)
		check(err)
		s += k + " " + v + "\n"
	}
	fmt.Fprintln(w, s)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
