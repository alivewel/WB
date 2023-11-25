package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello, world!"
	sep := ","
	before, after, found := strings.Cut(s, sep)
	fmt.Printf("Cut(%q, %q) = %q, %q, %v\n", s, sep, before, after, found)
}
