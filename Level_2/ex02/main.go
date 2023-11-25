package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func parserString(in string) string {
	out := make([]rune, 10)
	counter := 1
	for i := 1; i < len(in); i++ {
		if in[i-1] == in[i] {
			counter++
		} else if in[i-1] != in[i] {
			out = append(out, rune(in[i-1]))
			out = append(out, rune('0'+counter))
			counter = 1
		}
	}
	out = append(out, rune(in[len(in)-1]))
	out = append(out, rune('0'+counter))
	return string(out)
}

func main() {
	// str := "qwe\x04\x05"
	str := "qwe\x045"
	// str := "a11bc2d5e"
	// str := "a0bc2d5e"
	// str := "abcd"
	// str := "45"
	// str := ""
	newStr := ""
	currentDigit := ""
	var digit int
	var prevSym rune
	escape := false
	for _, char := range str {
		if escape {
			if unicode.IsDigit(char) {
				currentDigit += string(char)
				digit, _ = strconv.Atoi(currentDigit)
			} else {
				newStr += string(prevSym)
			}
			prevSym = char
		} else {
			if unicode.IsDigit(char) {
				currentDigit += string(char)
				digit, _ = strconv.Atoi(currentDigit)
			} else {
				for i := digit - 1; i > 0; i-- {
					newStr += string(prevSym)
				}
				if digit != 0 {
					newStr += string(prevSym)
				}
				currentDigit = ""
				digit = -1
				prevSym = char
			}
		}
		if char == '/' {
			escape = true
		}
	}
	newStr += string(prevSym)
	fmt.Println(newStr)
}
