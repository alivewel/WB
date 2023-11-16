package main

import "fmt"

func main() {
	str := "Строка String 😊" // Ваша строка

	for i, char := range str {
		fmt.Printf("%d %c\n", i, char)
	}
}
