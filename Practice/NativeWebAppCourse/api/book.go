package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}

//create sample books
var books = map[string]Book{
	"01234": Book{Title: "Cloud Native", Author: "Writer", ISBN: "01234"},
	"56789": Book{Title: "Test Book Two", Author: "Second Author", ISBN: "56789"},
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

func AllBooks() []Book {
	values := make([]Book, len(books))
	var index int
	for _, book := range books {
		values[index] = book
		index++
	}
	return values
}

//BookHandler
func BooksHandler(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case http.MethodGet:
		books := AllBooks()
		writeJSON(w, books)
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		book := FromJSON(body)
		isbn, created := CreateBook(book)
		if created {
			w.Header().Add("Location", "/api/books/"+isbn)
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusConflict)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request Method"))
	}
}

func CreateBook(b Book) (string, bool) {

}

func writeJSON
