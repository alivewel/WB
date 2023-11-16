package main

// import (
// 	// "fmt"
// )

func main() {
	x := getValue()
	println(*x/2)
}

func getValue() *int {
	x := 4
	return &x
}
