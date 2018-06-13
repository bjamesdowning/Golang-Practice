package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Name   string        `json:"name" bson:"name"`
	Gender string        `json:"gender" bson":gender"`
	Age    int           `json:"age" bson:"age"`
	Id     bson.ObjectId `json:"id" bson:"id"`
}

func main() {
	session, err := mgo.Dial("localhost:32768")
	if err != nil {
		fmt.Println("session error: ", err)
	}
	u := User{
		Name:   "John",
		Gender: "Female",
		Age:    22,
		Id:     bson.NewObjectId(),
	}

	v := User{}
	session.DB("go-mongo").C("users").Insert(u)
	fmt.Println("*************Getting User******************")

	if err = session.DB("go-mongo").C("users").FindId("5b21577a0e48dc396663b77d").One(&v); err != nil {
		fmt.Println(err)
	}

}
