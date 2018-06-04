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
	Role  string
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
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/signup", signUp)
	http.HandleFunc("/adminhome", adminHome)
	http.HandleFunc("/userhome", userHome)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)
	tmpl.ExecuteTemplate(w, "index.html", u)
}

func login(w http.ResponseWriter, r *http.Request) {
	if loggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		pword := r.FormValue("password")

		u, ok := dbUsers[email]
		if !ok {
			http.Error(w, "Username/Password Error", http.StatusForbidden)
			return
		}
		err := bcrypt.CompareHashAndPassword(u.Pword, []byte(pword))
		if err != nil {
			http.Error(w, "Username/Password Error", http.StatusForbidden)
			return
		}
		//Username and Password match. Create a session
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		dbSessions[c.Value] = email

		if u.Role == "admin" {
			http.Redirect(w, r, "/adminhome", http.StatusSeeOther)
			return
		} else if u.Role == "user" {
			http.Redirect(w, r, "/userhome", http.StatusSeeOther)
			return
		}
	}
	tmpl.ExecuteTemplate(w, "login.html", nil)
}

func logout(w http.ResponseWriter, r *http.Request) {
	if !loggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	c, _ := r.Cookie("session")
	//remove session from db
	delete(dbSessions, c.Value)
	//delete cookie by assigning new cookie with negative age value
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)
	http.Redirect(w, r, "/", http.StatusSeeOther)
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
		role := r.FormValue("role")
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
		u = user{email, fname, lname, role, bs}
		dbSessions[c.Value] = email
		dbUsers[email] = u
		if u.Role == "admin" {
			http.Redirect(w, r, "/adminhome", http.StatusSeeOther)
			return
		} else if u.Role == "user" {
			http.Redirect(w, r, "/userhome", http.StatusSeeOther)
			return
		}
	}

	tmpl.ExecuteTemplate(w, "signup.html", u)

}

func adminHome(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)
	if !loggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if u.Role != "admin" {
		http.Error(w, "Only Admins Allowed", http.StatusForbidden)
		return
	}
	tmpl.ExecuteTemplate(w, "admin.html", u)
}

func userHome(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)
	if !loggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if u.Role != "user" {
		http.Error(w, "Users Homepage", http.StatusForbidden)
		return
	}
	tmpl.ExecuteTemplate(w, "user.html", u)
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
