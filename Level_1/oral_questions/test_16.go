package main

import "fmt"

func appendToSliceByReference(s *[]int) {
	*s = append(*s, 4, 5)
}

func main() {
	slice := []int{1, 2, 3}
	appendToSliceByReference(&slice)

	fmt.Println("Измененный слайс:", slice)
}
