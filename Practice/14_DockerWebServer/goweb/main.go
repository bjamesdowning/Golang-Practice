package main

import (
	"html/template"
	"log"
	"net/http"
)

//Test webapp to turn into container
var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", root)
	http.ListenAndServe(":8080", nil)
}

func root(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	tmpl.ExecuteTemplate(w, "index.gohtml", r.Form)
}
