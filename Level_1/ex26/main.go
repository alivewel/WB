package main

import "fmt"

func checkString(str string) bool {
	checkMap := make(map[rune]struct{})
	for _, elem := range str {
		_, ok := checkMap[elem]
		if ok {
			return false
		}
		checkMap[elem] = struct{}{}
	}
	return true
}

func main() {
	str := "abcda"
	result := checkString(str)
	fmt.Println(result)
}
