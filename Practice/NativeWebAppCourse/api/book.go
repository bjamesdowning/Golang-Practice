package api

import (
	"encoding/json"
	"net/http"
)

type Book struct {
	Title  string
	Author string
	ISBN   string
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

//BookHandler
func BookHandler(w http.ResponseWriter, r *http.Request) {

}
