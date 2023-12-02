package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	str := "a4bc2d5e"
	// str := "qwe\\\\5"
	// str := "qwe\x04\x05"
	// str := "qwe\x045"
	// str := "a11bc2d5e"
	// str := "a0bc2d5e"
	// str := "abcd"
	// str := "45"
	// str := ""
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
				// fmt.Println("currentDigit", currentDigit)
				// fmt.Println("digit", digit)
			}
		} else if unicode.IsLetter(char) {
			if digit == 0 {
				tempStr += string(char)
				fmt.Println("tempStr2", tempStr)
			} else {
				for i := digit - 1; i > 0; i-- {
					// newStr += string(char)
					newStr += tempStr
				}
				if digit != 0 {
					// newStr += string(char)
					newStr += tempStr
				}
				tempStr += string(char)
				newStr += tempStr
				fmt.Println("tempStr", tempStr)
				currentDigit = ""
				tempStr = ""
				digit = -1
			}
		} else if char == '\\' {
			countSlash++
			fmt.Println("countSlash", countSlash)
			if countSlash == 2 {
				newStr += string(char)
				countSlash = 0
				fmt.Println("!", string(char))
			}
		}
	}
	fmt.Println(newStr)
}
