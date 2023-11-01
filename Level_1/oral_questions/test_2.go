package main

import "fmt"

func main() {
	slice := []int{}
	slice = append(slice, 0)
	slice = append(slice, 1)
	// for i := 0; i < len(slice); i++ {
	// // fmt.Println(slice[i])
	// fmt.Println(len(slice))
	// 	slice = append(slice, i+1)
	// }
	for i := range slice {
		// fmt.Println(slice[i])
		fmt.Println(len(slice))
		slice = append(slice, i+1)
	}
}
