package main

import "fmt"

func main() {
	// Создаем пустой слайс целых чисел
	var numbers []int

	// Добавляем элементы в слайс с использованием append
	numbers = append(numbers, 1)
	numbers = append(numbers, 2, 3)
	// numbers = append(numbers, 4, 5)
	append(numbers, 4, 5)

	fmt.Println(numbers)
}
