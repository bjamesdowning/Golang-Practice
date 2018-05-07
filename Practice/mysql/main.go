package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//Program following book https://goo.gl/wepCxp
//If using docker container running mysql, be sure to use 5.7
//Latest container image pulls mysql 8 which throws an error about
//Not supporting the authentication plugin. I'm not sure why.
//docker command: docker run -p 3306:3306 --name go-mysql -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=test -d mysql:5.7
func main() {
	db, err := sql.Open("mysql", "root:password@tcp(:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(
		"CREATE TABLE IF NOT EXISTS test.hello(world varchar(50))")
	if err != nil {
		log.Fatal(err)
	}

	res, err := db.Exec(
		"INSERT INTO test.hello(world) VALUES('hello world!')")
	if err != nil {
		log.Fatal(err)
	}
	rowCount, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("inserted %d rows", rowCount)

	rows, err := db.Query("SELECT * FROM test.hello")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var s string
		err = rows.Scan(&s)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("found row containing %q", s)
	}
	rows.Close()
}
