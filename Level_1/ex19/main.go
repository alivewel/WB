package main

import "fmt"

func reverseString(str string) string {
	runes := []rune(str)
	reverseStr := make([]rune, len(runes))
	for i := len(runes) - 1; i >= 0; i-- {
		reverseStr = append(reverseStr, runes[i])
	}
	return string(reverseStr)
}

func main() {
	str := "Привет, 🌍! 你好, 世界!"
	fmt.Println(str)
	result := reverseString(str)
	fmt.Println(result)
}
