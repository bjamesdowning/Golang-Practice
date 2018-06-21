package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//User is a struct
type User struct {
	Name   string        `json:"name" bson:"name"`
	Gender string        `json:"gender" bson:"gender"`
	Age    int           `json:"age" bson:"age"`
	ID     bson.ObjectId `json:"_id" bson:"_id"`
}

var ErrNotFound error

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		fmt.Println("session error: ", err)
	}
	defer session.Close()

	u := User{
		Name:   "Dexter1",
		Gender: "Male",
		Age:    55,
		ID:     bson.NewObjectId(),
	}

	c := session.DB("go-mongo").C("users")
	err = c.Insert(u)
	if err != nil {
		fmt.Println("Could not add user:", err)
	}

	v := User{}
	err = c.Find(bson.M{"_id": u.ID}).One(&v)
	if err != nil {
		fmt.Println("Server Error Finding User:", err)
	}
	fmt.Println(v)

	newUser := User{
		Name:   "Notnew",
		Gender: "Male",
		Age:    55,
		ID:     bson.NewObjectId(),
	}

	testUser := User{}
	//compare new user to table, if already exists do not add
	err = c.Find(bson.M{"name": "Notnew"}).One(&testUser)
	if err == mgo.ErrNotFound {
		err = c.Insert(newUser)
		if err != nil {
			fmt.Println("Could not add user:", err)
		}
	} else {
		fmt.Println("User Already Exists")
	}
}
