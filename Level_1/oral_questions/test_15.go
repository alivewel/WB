package main

import "fmt"

func appendToSliceByValue(s []int) []int {
	s = append(s, 4, 5)
	return s
}

func main() {
	slice := []int{1, 2, 3}
	modifiedSlice := appendToSliceByValue(slice)

	fmt.Println("Оригинальный слайс:", slice)
	fmt.Println("Измененный слайс:", modifiedSlice)
}
