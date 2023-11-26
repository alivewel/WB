package main

import (
	"fmt"
)

func generateRange(start, offset, maxNum int, result []int) []int {
	if offset > 0 {
		// В случае положительного offset
		for i := start + 1; i <= start+offset; i++ {
			if i > maxNum {
				break
			}
			result = append(result, i)
		}
	} else if offset < 0 {
		// В случае отрицательного offset
		for i := start - 1; i >= start+offset; i-- {
			if i < 0 {
				break
			}
			result = append(result, i)
		}
	} else {
		// В случае offset равного нулю
		result = append(result, start)
	}
	// if true {
	// 	result = generateRange(start, -offset, maxNum, result)
	// }
	return result
}

func main() {
	var result []int
	start := 5
	offset := -3
	maxNum := 7

	result = generateRange(start, offset, maxNum, result)

	start = 5
	offset = 3
	maxNum = 7

	result = generateRange(start, offset, maxNum, result)
	fmt.Println(result)
}
