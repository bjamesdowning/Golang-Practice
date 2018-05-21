package api

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestBookToJSON(t *testing.T) {
	book := Book{
		Title: "Cloud Native",
		Author: "Writer",
		ISBN: "1243532",
	}
	json := book.ToJSON()
	
	assert.Equal(t, `{"Title":"Cloud Native","Author","Writer","ISBN","1243532"}`),
		string(json), "Book JSON Marshalling wrong") 
}