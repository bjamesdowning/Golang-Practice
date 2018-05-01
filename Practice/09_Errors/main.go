package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var ErrMath = errors.New("Cannot find square root of negative number")

//init function runs before main, allows setup of program flags or environment
func init() {
	f, err := os.Create("Log.txt")
	if err != nil {
		fmt.Println(err)
	}
	//altering where log.Println will send the err message, rather than stdou, sends to this write interface, f
	log.SetOutput(f)
}

func main() {
	//throway var, used to simulate error.
	_, err := sqrt(-10)
	if err != nil {
		log.Println("Error: ", err) //log.Println similar to fmt.Prinlnt but sends log to wherever you've setup the output. Or, just sends to stdout
		// log.Fataln("Error: ", err) //same as log.Println, but also calls os.Exit to kill program
		//panic(err) //calls panic after writing the log message. Pulled from builtin package
	}
}

func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0, ErrMath
	}
	//just a test case
	return 42, nil
}
