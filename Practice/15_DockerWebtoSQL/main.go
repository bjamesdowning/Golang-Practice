package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

//practice application to run a generic front end to add data to a mysql backend.
//goal to practice using Docker Compose to have MySQL container and FE app deploy as a service.
var tmpl *template.Template
var db *sql.DB

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {

	mux := httprouter.New()
	mux.GET("/", welcome)
	mux.POST("/insert", insert)
	mux.GET("/read", read)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func welcome(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.ExecuteTemplate(w, "index.gohtml", r.Form)
	if err != nil {
		log.Fatal(err)
	}
}

func insert(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db, err := sql.Open("mysql", "root:password@tcp(:3306)/test")
	if err != nil {
		log.Fatal("DB Connect Error: ", err)
	}
	defer db.Close()

	_, err = db.Exec(
		"CREATE TABLE IF NOT EXISTS test.table(column_one varchar(50), column_two varchar(50))")
	if err != nil {
		log.Fatal(err)
	}

	err = r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	k := strings.Join(r.PostForm["key"], "")
	v := strings.Join(r.PostForm["value"], "")
	res, err := db.Exec(
		"INSERT INTO test.table(column_one, column_two) VALUES('" + k + "','" + v + "')")
	if err != nil {
		log.Fatal(err)
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "inserted %d rows", rowCount)
}

func read(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db, err := sql.Open("mysql", "root:password@tcp(:3306)/test")
	if err != nil {
		log.Fatal("DB Connect Error: ", err)
	}
	defer db.Close()

	rows, err := db.Query(`SELECT * FROM test.table;`)
	if err != nil {
		log.Fatal(err)
	}

	// data to be used in query
	var s, k, v string
	s = "RETRIEVED RECORDS:\n"

	// query
	for rows.Next() {
		err = rows.Scan(&k, &v)
		if err != nil {
			log.Fatal(err)
		}
		s += k + " " + v + "\n"
	}
	fmt.Fprintln(w, s)
}
