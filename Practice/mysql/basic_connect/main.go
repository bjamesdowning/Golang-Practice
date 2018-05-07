package main

import (
	"database/sql"
	"flag"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//Program following book https://goo.gl/wepCxp
//If using docker container running mysql, be sure to use 5.7
//Latest container image pulls mysql 8 which throws an error about
//Not supporting the authentication plugin. I'm not sure why.
//docker command: docker run -p 3306:3306 --name go-mysql -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=test -d mysql:5.7
func main() {
	//practice using flags. This flag allows you to pass in a value to be added to the test table.
	key := flag.String("k", "KEY", "a key is added to table: CMDB")
	value := flag.String("v", "VALUE", "a value added to table: CMDB")
	flag.Parse()

	db, err := sql.Open("mysql", "root:password@tcp(:3306)/test")
	check(err)

	defer db.Close()

	_, err = db.Exec(
		"CREATE TABLE IF NOT EXISTS test.CMDB(RouterID varchar(50), IP_Addr varchar(16))")
	check(err)

	res, err := db.Exec(
		"INSERT INTO test.CMDB(RouterID, IP_Addr) VALUES('" + *key + "','" + *value + "')")
	check(err)

	rowCount, err := res.RowsAffected()
	check(err)

	log.Printf("inserted %d rows", rowCount)

	rows, err := db.Query("SELECT * FROM test.CMDB")
	check(err)

	for rows.Next() {
		var k, v string
		err = rows.Scan(&k, &v)
		check(err)
		log.Printf("found row containing %q %q", k, v)
	}
	rows.Close()
}

//function to check errors without rewriting same code repeatedly
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
