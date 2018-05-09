package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//learning structs

type IP struct {
	Oct1 uint8
	Oct2 uint8
	Oct3 uint8
	Oct4 uint8
}

type RTR struct {
	Hname string `json:"hostname"`
	Uname string `json:"username"`
	Pword string `json:"password"`
	IP
}

//not good.
func (IP IP) dotNotation() (uint8, string, uint8, string, uint8, string, uint8) {
	return IP.Oct1, ".", IP.Oct2, ".", IP.Oct3, ".", IP.Oct4
}

func main() {
	obj1 := RTR{
		Hname: "R1",
		Uname: "cisco",
		Pword: "cisco",
		IP: IP{
			Oct1: 192,
			Oct2: 168,
			Oct3: 1,
			Oct4: 5,
		},
	}
	//obj2 := foo{"Sara", "Dowell", "", 37}
	fmt.Println(obj1.IP.dotNotation())
	fmt.Println(obj1)
	bs, _ := json.Marshal(obj1)
	fmt.Println(bs)
	fmt.Printf("%T, \n", bs)
	fmt.Println(string(bs))
	fmt.Println("---------")
	json.NewEncoder(os.Stdout).Encode(&obj1)
	//fmt.Println(obj2.fname, obj2.age)
}
