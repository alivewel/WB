package main

import (
	"fmt"
	"sort"
)

var (
	flagA bool = false
	flagB bool = false
	flagC bool = true
)

func generateRange(start, offset, maxNum int, result *[]int) {
	if offset > 0 {
		// В случае положительного offset
		for i := start + 1; i <= start+offset; i++ {
			if i > maxNum {
				break
			}
			*result = append(*result, i)
		}
	} else if offset < 0 {
		// В случае отрицательного offset
		for i := start - 1; i >= start+offset; i-- {
			if i < 0 {
				break
			}
			*result = append(*result, i)
		}
	} else {
		// В случае offset равного нулю
		*result = append(*result, start)
	}
}

func addArr(start, offset, maxNum int, result *[]int) {
	if flagA || flagC {
		generateRange(start, offset, maxNum, result)
	}
	if flagB || flagC {
		generateRange(start, -offset, maxNum, result)
	}
	*result = append(*result, start)
}

func removeDuplicates(input []int) []int {
	uniqueMap := make(map[int]bool)
	var result []int
	for _, num := range input {
		if !uniqueMap[num] {
			// Если элемент не встречался ранее, добавляем его в результат и карту
			result = append(result, num)
			uniqueMap[num] = true
		}
	}
	return result
}

func main() {
	var result []int
	start := 5
	offset := -3
	maxNum := 7

	addArr(start, offset, maxNum, &result)
	result = removeDuplicates(result)
	sort.Ints(result)
	// generateRange(start, offset, maxNum, &result)

	// start = 5
	// offset = 3
	// maxNum = 7

	// generateRange(start, offset, maxNum, &result)
	fmt.Println(result)
}
