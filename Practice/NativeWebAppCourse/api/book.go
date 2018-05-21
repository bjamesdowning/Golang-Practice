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

func (b Book) ToJSON() []byte {
	ToJSON, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return ToJSON
}

func FromJSON(d []byte) Book {
	return Book{}
}

func BookHandler(w http.ResponseWriter, r *http.Request) {

}
