package main

import (
	"fmt"
)

// in: AAABBBCCDDDDAAA
// out: A3B3C2D3A3

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
	out := parserString("ABBBCCDDDDAAA")
	fmt.Println(out)
}
