package main

import (
	"fmt"
	"strings"
)

func reverseString(inputStr string) string {
	strArr := strings.Fields(inputStr)
	lenStr := len(strArr)
	revArr := make([]string, lenStr)
	for i, j := 0, lenStr-1; i < lenStr; i, j = i+1, j-1 {
		revArr[i] = strArr[j]
	}
	return strings.Join(revArr, " ")
}

func main() {
	str := "snow dog sun"
	revStr := reverseString(str)
	fmt.Println(revStr)
}
