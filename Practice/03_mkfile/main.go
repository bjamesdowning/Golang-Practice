package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer file.Close()
	file.WriteString("TEST \n TEST \n TEST")
}
