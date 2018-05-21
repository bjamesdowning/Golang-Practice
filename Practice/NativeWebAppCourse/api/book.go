package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//Book type
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

//ToJSON for marshaling
func (b Book) ToJSON() []byte {
	ToJSON, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return ToJSON
}

//FromJSON for unmarshaling
func FromJSON(d []byte) Book {
	book := Book{}
	err := json.Unmarshal(d, &book)
	if err != nil {
		panic(err)
	}
	return book
}

//AllBooks retrieves all books
func AllBooks() []Book {
	values := make([]Book, len(books))
	var index int
	for _, book := range books {
		values[index] = book
		index++
	}
	return values
}

//BooksHandler acts as handler for book endpoint for all books
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

//BookHandler function takes care of single book requests
func BookHandler(w http.ResponseWriter, r *http.Request) {
	isbn := r.URL.Path[len("/api/books/"):]

	switch method := r.Method; method {
	case http.MethodGet:
		book, found := GetBook(isbn)
		if found {
			writeJSON(w, book)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodPut:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		book := FromJSON(body)
		exists := UpdateBook(isbn, book)
		if exists {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodDelete:
		DeleteBook(isbn)
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request Method"))
	}
}

//CreateBook creates a new book if it doesn't exist
func CreateBook(b Book) (string, bool) {
	_, exists := books[b.ISBN]
	if exists {
		return "", false
	}
	books[b.ISBN] = b
	return b.ISBN, true
}

func writeJSON(w http.ResponseWriter, i interface{}) {
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}

//GetBook finds book in map based on key of isbn
func GetBook(isbn string) (Book, bool) {
	book, found := books[isbn]
	return book, found
}

//UpdateBook edits a book within the map
func UpdateBook(isbn string, b Book) bool {
	_, exists := books[isbn]
	if exists {
		books[isbn] = b
	}
	return exists
}

//DeleteBook removes book from map
func DeleteBook(isbn string) {
	delete(books, isbn)
}
