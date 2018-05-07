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
//test table by created from origination
//docker command:
//docker run -p 3306:3306 --name go-mysql -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=test -d mysql:5.7
func main() {
	db, err = sql.Open("mysql", "root:password@tcp(:3306)/test")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", readAll)
	err := http.ListenAndServe(":8080", nil)
	check(err)
}

func readAll(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`SELECT * FROM test.hello;`)
	check(err)

	// data to be used in query
	var s, name string
	s = "RETRIEVED RECORDS:\n"

	// query
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		s += name + "\n"
	}
	fmt.Fprintln(w, s)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
