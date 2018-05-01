package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

//create a map object to take in the json data
//use open interface to allow various input types
var obj map[string]interface{}

func main() {
	f, err := os.Open("data.json")
	if err != nil {
		log.Println("Error: ", err)
	}
	defer f.Close()

	json.NewDecoder(f).Decode(&obj) //Ensure Decode/Encode takes in a pointer
	fmt.Println(obj)
}
