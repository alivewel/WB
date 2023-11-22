package main

import "fmt"

func main() {
	myMap := make(map[int]int)
	myMap[1] = 42

	// Получаем адрес значения из map
	addr := &myMap[1]

	// Выводим значение, на которое указывает адрес
	fmt.Println(*addr)
}
