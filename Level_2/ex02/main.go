package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	str := "a4bc2d5e"
	result := expandString(str)
	fmt.Println(result)
}

func expandString(str string) string {
	newStr, currentDigit := "", ""
	digit, countSlash := 0, 0
	tempStr := ""
	for _, char := range str {
		if unicode.IsDigit(char) {
			if countSlash == 1 {
				newStr += string(char)
			} else {
				currentDigit += string(char)
				digit, _ = strconv.Atoi(currentDigit)
			}
		} else if unicode.IsLetter(char) {
			if digit == 0 {
				tempStr += string(char)
			} else {
				for i := digit - 1; i > 0; i-- {
					newStr += tempStr
				}
				if digit != 0 {
					newStr += tempStr
				}
				tempStr += string(char)
				newStr += tempStr
				currentDigit = ""
				tempStr = ""
				digit = -1
			}
		} else if char == '\\' {
			countSlash++
			if countSlash == 2 {
				newStr += string(char)
				countSlash = 0
			}
		}
	}
	return newStr
}
