package main

import (
	"fmt"

	"github.com/ttacon/chalk"
)

//quick practice for importing using 'go get' and using 3rd part package.

func main() {
	fmt.Println(chalk.Red, "red text", chalk.ResetColor)
}
