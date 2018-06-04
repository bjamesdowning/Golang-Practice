package main

import (
	"html/template"
	"net/http"

	"github.com/satori/go.uuid"
)

type user struct {
	Email string
	Fname string
	Lname string
}

var tmpl *template.Template
var dbSessions = map[string]string{}
var dbUsers = map[string]user{}

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/verify", verify)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	//Check for cookie
	c, err := r.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4() //throw away error.
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}
	//At this point, the user has a cookie
	var u user
	if email, ok := dbSessions[c.Value]; ok {
		u = dbUsers[email]
	}
	//process form to create user
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		fname := r.FormValue("firstname")
		lname := r.FormValue("lastname")
		u = user{email, fname, lname}
		dbSessions[c.Value] = email
		dbUsers[email] = u
	}

	tmpl.ExecuteTemplate(w, "index.html", u)
}

func verify(w http.ResponseWriter, r *http.Request) {
	//Grab cookie. If no cookie, redirect to index
	c, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	email, ok := dbSessions[c.Value]
	if !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	u := dbUsers[email]
	tmpl.ExecuteTemplate(w, "verify.html", u)
}
