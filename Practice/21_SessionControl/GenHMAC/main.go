package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {
	//generate hmac based on input.
	c := getCode("test@email.com")
	fmt.Println(c)
	//show each time the same input will create same output
	c = getCode("test@email.com")
	fmt.Println(c)
}

func getCode(s string) string {
	h := hmac.New(sha256.New, []byte("ourkey"))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}
