package main

import "fmt"

func main() {
	set := make(map[string]int)

	elements := []string{"cat", "cat", "dog", "cat", "tree"}

	for _, element := range elements {
		set[element]++
	}

	// Вывод множества с повторами и их количеством
	// for element, count := range set {
	// 	fmt.Printf("%s: %d\n", element, count)
	// }

	for element, count := range set {
		for i := 0; i < count; i++ {
			fmt.Println(element)
		}
	}
}
