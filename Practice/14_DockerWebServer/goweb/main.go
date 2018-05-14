package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//Test webapp to turn into container using third party mux
var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	mux := httprouter.New() //create a new mux as opposed to using DefaultServeMux
	mux.GET("/", root)
	http.ListenAndServe(":8080", mux) //pass in handler, mux

	//below is using default servemux
	//http.HandleFunc("/", root)
	//http.ListenAndServe(":8080", nil)
}

//with this mux, need to add the 3rd parameter for vars
func root(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	tmpl.ExecuteTemplate(w, "index.gohtml", r.Form)
	if err != nil {
		log.Fatal(err)
	}
}
