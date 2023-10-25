package main

import (
	"fmt"
	"strconv"
)

// in: AAABBBCCDDDDAAA
// out: A3B3C2D3A3

func parserString(in string) string {
	out := ""
	prev := '0'
	counter := 1
	for i, char := range in {
		if char == prev {
			counter++
		} else if char != prev && i != 0 {
			out += string(prev)
			out += strconv.Itoa(counter)
			counter = 1
		}
		prev = char
	}
	out += string(prev)
	out += strconv.Itoa(counter)
	return out
}

func main() {
	out := parserString("AAABBBCCDDDDAAA")
	fmt.Println(out)
}
