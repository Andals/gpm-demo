package main

import (
	"andals/gobox/shell"
	"fmt"
)

func main() {
	result := shell.RunCmd("ls -al")

	fmt.Println(result)
}
