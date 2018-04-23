package main

import "fmt"

//learning structs

type ip struct {
	oct1 uint8
	oct2 uint8
	oct3 uint8
	oct4 uint8
}

type rtr struct {
	hname string
	uname string
	pword string
	ip
}

//not good.
func (ip ip) dotNotation() (uint8, string, uint8, string, uint8, string, uint8) {
	return ip.oct1, ".", ip.oct2, ".", ip.oct3, ".", ip.oct4
}

func main() {
	obj1 := rtr{
		hname: "R1",
		uname: "cisco",
		pword: "cisco",
		ip: ip{
			oct1: 192,
			oct2: 168,
			oct3: 1,
			oct4: 5,
		},
	}
	//obj2 := foo{"Sara", "Dowell", "", 37}
	fmt.Println(obj1.ip.dotNotation())
	fmt.Println(obj1)
	//fmt.Println(obj2.fname, obj2.age)
}
