package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(1)
	for num, arg := range os.Args[1:] {
		fmt.Println(num, arg)
	}
}
