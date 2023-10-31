package main

import (
	"fmt"
)

func someAction(v []int8, b int8) {
	v[0] = 100       // изменение сохранится
	v = append(v, b) // записали значение, новый массив вернули в переменную v, но никуда не сохранили его и не записали
}

func main() {
	var a = []int8{1, 2, 3, 4, 5}
	someAction(a, 6)
	fmt.Println(a)
}
