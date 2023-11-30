package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	str := "qwe\\\\5"
	// str := "qwe\x04\x05"
	// str := "qwe\x045"
	// str := "a11bc2d5e"
	// str := "a0bc2d5e"
	// str := "abcd"
	// str := "45"
	// str := ""
	newStr := ""
	currentDigit := ""
	// var digit int
	// var countSlash int
	digit, countSlash := 0, 0
	// var prevSym rune
	for i, char := range str {
		if unicode.IsDigit(char) {
			currentDigit += string(char)
			digit, _ = strconv.Atoi(currentDigit)
		} else if unicode.IsLetter(char) {
			if i == 0 {
				newStr += string(char)
			}
			for i := digit - 1; i > 0; i-- {
				// newStr += string(prevSym)
				newStr += string(char)
			}
			if digit != 0 {
				// newStr += string(prevSym)
				newStr += string(char)
			}
			currentDigit = ""
			digit = -1
			// prevSym = char
		} else if char == '\\' {
			countSlash++
			fmt.Println("countSlash", countSlash)
			if countSlash == 2 {
				// if prevSym != '0' {
				// 	newStr += string(prevSym)
				// }
				newStr += string(char)
				countSlash = 0
				fmt.Println("!", string(char))
			}
		}
	}
	fmt.Println(newStr)
}
