package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

//Msg struct used to encode/decode json
type Msg struct {
	Fname string `json:"fname"`
	Lname string `json:"lname"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func main() {
	msg := []Msg{
		{"John", "Doe", "Jdoe@email.com", 27},
		{"Dana", "Sumthin", "Dsumthin@email.com", 27},
		{"Jerry", "Else", "JElse@email.com", 27},
		{"Phil", "Other", "Pother@email.com", 27},
	}

	createJSON(msg)

	f, err := os.Open("data.json")
	if err != nil {
		log.Println("Error: ", err)
	}
	defer f.Close()

	msg2 := []Msg{}
	dec := json.NewDecoder(f)
	dec.Decode(&msg2)
	fmt.Println(msg2)
}

func createJSON(msg []Msg) {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.Encode(msg)

	f, err := os.Create("data.json")
	if err != nil {
		log.Println("Error: ", err)
	}
	defer f.Close()
	io.Copy(f, buf)
}
