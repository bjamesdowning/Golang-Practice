package main

import (
	"fmt"

	"github.com/bjamesdowning/GoLearning/Practice/01_hello/hellopkg"
)

var x = 42 //This is declaring 'x' type int and assigned value 42
//This is PACKAGE level scope, meaning every func can use it

func main() {
	var msg string
	a := 42 //same as 'x' but only accessbile within this block
	fmt.Println(a)
	fmt.Print("Enter Message to Echo: ")
	fmt.Scan(&msg)
	hellopkg.Hello(msg)
	new := test{15, 26}
	new.add()
}

/* declare variable
*	Shorthand: within a func, simple > a := 10, or a := "golang"
* DOUBLE quotes matter. SIde now, caitlization also matters
* lowercase = not visible to outside package, captials mean = visible
* to outside package
 */
