package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

//Writing a cookie to track amount of times a user goes to site.
func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	//checking if a cookie exists. If not, create one.
	cookie, err := r.Cookie("mycookie")
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "mycookie",
			Value: "0",
		}
	}
	//Now cookie is created. Convert the string Value to integer
	count, err := strconv.Atoi(cookie.Value)
	if err != nil {
		log.Fatalln(err)
	}
	//Add 1 to the count (which is now an integer = to cookie.Value
	count++
	//Convert cookie.Value back to a string
	cookie.Value = strconv.Itoa(count)
	//rewrite the cookie with the new value
	http.SetCookie(w, cookie)
	//this simply writes the value to the browser
	io.WriteString(w, cookie.Value)
}
