package main

import "fmt"

func main() {
	m := map[string]int{
		"key 1": 12,
		"key 2": 23,
		"key 3": 100,
	}

	m2 := make(map[int]int)
	fmt.Println(m2 == nil)

	if val, ok := m["key 2"]; ok { //comma ok idion
		fmt.Println(val)
	} else {
		fmt.Println("no value")
	}

	for key, val := range m { //loop over map
		fmt.Println(key, "--", val)
	}
	fmt.Println(len(m))
}
