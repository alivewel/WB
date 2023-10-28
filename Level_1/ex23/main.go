package main

import (
	"fmt"
	"reflect"
)

func removeElemStr(slice interface{}) interface{} {
	sliceType := reflect.TypeOf(slice)
	newSlice := reflect.New(sliceType).Elem()
	return newSlice.Interface()
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}

	// a = append(a[0:2], a[3:]...) // элемент в середине

	b := make([]int, 0, len(a)-1)
	// a = append(b, a[1:]...) // первый элемент
	a = append(b, a[0:(len(a)-1)]...) // последний элемент

	fmt.Println(a) // [1 2 4 5 6 7]
}
