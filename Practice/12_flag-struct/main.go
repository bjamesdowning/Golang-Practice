package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

//IP is a struct to hold IPs
type IP struct {
	oct1 int
	oct2 int
	oct3 int
	oct4 int
}

func main() {
	inIP := flag.String("i", "", "Define IP")
	inSUB := flag.String("s", "", "Define subnet mask")
	flag.Parse()
	ip := strings.Split(*inIP, ".")
	sub, _ := strconv.Atoi(*inSUB)
	addr := make([]int, 4)
	for i, x := range ip {
		addr[i], _ = strconv.Atoi(x)
	}
	new := IP{
		oct1: addr[0],
		oct2: addr[1],
		oct3: addr[2],
		oct4: addr[3],
	}
	fmt.Printf("IP address %v/%v\n. In binary %b", new, sub, new)
}
