package main

import (
	"fmt"
)

func main() {
	strings := []string{"This ", "is ", "even ",
		"more ", "performant "}

	bs := make([]byte, 100)
	bl := 0

	for _, val := range strings {
		bl += copy(bs[bl:], []byte(val))
		fmt.Println(bl)
	}

	fmt.Println(string(bs[:]))
}

// bl += copy(bs[bl:], []byte(val)): 
// В этой строке выполняется копирование байтов из строки val в срез bs
//  начиная с текущей позиции bl. 
// Функция copy возвращает количество скопированных байтов, которое добавляется к bl.