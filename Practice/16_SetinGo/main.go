package main

import "fmt"

//represent a SET

type set map[string]struct{}

//creates type Set. a Set is a structure which stores values without allowing a value to be repeated
//use the keys and strings, and the values as empty structs. Use emtpy structs use no memory. So, this
//will create a map with just keys, where the keys cannot match

func main() {
	s := make(set)
	s["item1"] = struct{}{}
	s["item2"] = struct{}{}
	s["item1"] = struct{}{} //Won't be added. Matches already added key.
	fmt.Println(getSetValues(s))
}

//this is just to retrieve only the keys for the map, forcing it to act like a set
func getSetValues(s set) []string {
	var retVal []string
	for k := range s {
		retVal = append(retVal, k)
	}
	return retVal
}
