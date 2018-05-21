package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/bjamesdowning/Practice/NativeWebAppCourse/api"
)

//basic echo web server. Allows environment variable PORT to dicate listening port.
//user 'export PORT=<port> to set
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/api/echo", echo)
	http.HandleFunc("/api/books", api.BookHandler)
	http.ListenAndServe(port(), nil)
}

//dynamic listening port
func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = ":8080"
	}
	return ":" + port
}

//responds with http code 200 and message
func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Cloud native go course on %s", os.Getenv("PORT"))
}

//echos query sent in URL, as in "<server:port>/api/echo?message=Some+Message+here"
func echo(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query()["message"][0]

	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, message)
}
