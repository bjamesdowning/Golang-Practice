package main

import (
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		return
	}
	defer file.Close()
	file.WriteString("TEST \n TEST \n TEST")
}
