package main

import "fmt"

func invertIndex(allIndexStr [][]int, allScanStr [][]string) [][]int {
	if true {
		result := make([][]int, len(allScanStr))
		for i := 0; i < len(allScanStr); i++ {
			for j := 0; j < len(allScanStr[i]); j++ {
				if !containsElemArr(allIndexStr[i], j) {
					result[i] = append(result[i], j)
				}
			}
		}
		return result
	}
	return nil
}

func containsElemArr(nums []int, target int) bool {
	for _, num := range nums {
		if num == target {
			return true
		}
	}
	return false
}

func main() {
	originalArray := [][]int{{3, 5, 7}, {1, 2, 3}}
	excludedNumbers := [][]string{{"ac", "ac", "ac", "ac", "ac"}, {"ab", "ac", "ac", "ac", "ac", "ac", "ac"}}

	fmt.Println("Original Array:", originalArray)
	originalArray = invertIndex(originalArray, excludedNumbers)

	fmt.Println("Excluded Numbers:", excludedNumbers)
	fmt.Println("Inverted Array:", originalArray)
}
