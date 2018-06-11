package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		fmt.Println("session error: ", err)
	}
	c := session.DB("testDB1").C("testCOL1")
	fmt.Println(c)
	fmt.Println(c.Find(nil).Count())

}
