package main

import (
	"fmt"
)

func main() {
	n := 0
	if true {
		n := 1
		n++
	}
	fmt.Println(n)
}
