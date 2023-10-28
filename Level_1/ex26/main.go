package main

import (
	"fmt"
	"strings"
)

func checkString(str string) bool {
	lowStr := strings.ToLower(str)
	checkMap := make(map[rune]struct{})
	for _, elem := range lowStr {
		_, ok := checkMap[elem]
		if ok {
			return false
		}
		checkMap[elem] = struct{}{}
	}
	return true
}

func main() {
	str := "abcdA"
	result := checkString(str)
	fmt.Println(result)
}
