package main

import (
	"html/template"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/satori/go.uuid"
)

type user struct {
	Email string
	Fname string
	Lname string
	Pword []byte
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
	http.HandleFunc("/signup", signUp)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)
	tmpl.ExecuteTemplate(w, "index.html", u)
}

func signUp(w http.ResponseWriter, r *http.Request) {
	if loggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	var u user
	//process form to create user
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		fname := r.FormValue("firstname")
		lname := r.FormValue("lastname")
		pword := r.FormValue("password")
		//Check if username is already in the db
		if _, ok := dbUsers[email]; ok {
			http.Error(w, "Username Taken", http.StatusForbidden)
			return
		}
		//Create a session
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		dbSessions[c.Value] = email
		//Store user info. First, encrpt pword
		bs, err := bcrypt.GenerateFromPassword([]byte(pword), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Server Error, Password Related", http.StatusInternalServerError)
			return
		}
		u = user{email, fname, lname, bs}
		dbSessions[c.Value] = email
		dbUsers[email] = u

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tmpl.ExecuteTemplate(w, "signup.html", u)

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

func getUser(w http.ResponseWriter, r *http.Request) user {
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
	return u
}

func loggedIn(r *http.Request) bool {
	c, err := r.Cookie("session")
	if err != nil {
		return false
	}
	email := dbSessions[c.Value]
	_, ok := dbUsers[email]
	return ok
}
