package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "GeeksforGeeks is a computer science portal !"
	v := strings.Fields(s)
	for i, j := range v {
		fmt.Printf("%d %s\n", i, j)
	}
	newArr := make([]string, len(v))
	for i, j := 0, len(v)-1; i < len(v); i, j = i+1, j-1 {
		newArr[i] = v[j]
	}

	fmt.Println(newArr)
}
