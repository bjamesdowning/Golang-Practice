package api

import (
	"encoding/json"
	"net/http"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}

//ToJSON
func (b Book) ToJSON() []byte {
	ToJSON, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return ToJSON
}

//FromJSON
func FromJSON(d []byte) Book {
	book := Book{}
	err := json.Unmarshal(d, &book)
	if err != nil {
		panic(err)
	}
	return book
}

//create sample books
var Books = []Book{
	Book{Title: "Cloud Native", Author: "Writer", ISBN: "1243532"},
	Book{Title: "Test Book Two", Author: "Second Author", ISBN: "686858484"},
}

//BookHandler
func BookHandler(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(Books)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}
