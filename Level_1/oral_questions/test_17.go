package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	elem := 4

	// Добавление элемента elem в срез slice
	newSlice := append(slice, elem)

	fmt.Println("Исходный срез:")
	fmt.Println(slice)

	fmt.Println("Новый срез:")
	fmt.Println(newSlice)
}
