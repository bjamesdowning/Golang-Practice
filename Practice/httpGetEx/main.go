package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	r, err := http.Get("http://api.theysaidso.com/qod.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer r.Body.Close()

	c, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(c))
}
