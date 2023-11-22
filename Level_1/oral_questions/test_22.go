package main

import "fmt"

type MyStruct struct {
	Field string
}

func main() {
	var myStruct *MyStruct

	// Попытка обращения к полю объекта, значение которого nil
	fmt.Println(myStruct.Field) // Это вызовет ошибку времени выполнения
}
