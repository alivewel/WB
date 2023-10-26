package main

import "fmt"

func checkType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("i is an int: %d\n", v)
	case string:
		fmt.Printf("i is a string: %s\n", v)
	case bool:
		fmt.Printf("i is a bool: %t\n", v)
	case chan interface{}:
		fmt.Println("i is a channel")
	case float64:
		fmt.Printf("i is a float64: %v\n", v)
	case rune:
		fmt.Printf("i is a rune: %c\n", v)
	case nil:
		fmt.Println("i is nil")
	default:
		fmt.Println("i is of an unknown type")
	}
}

func main() {
	var i interface{}

	i = 123
	checkType(i)

	i = "Hello, WB!"
	checkType(i)

	i = true
	checkType(i)

	ch := make(chan interface{})
	i = ch
	checkType(i)

	i = 1.23
	checkType(i)

	i = 'E'
	checkType(i)

	i = nil
	checkType(i)
}
