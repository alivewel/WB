package main

import (
	"errors"
	"fmt"
)

func getBit(num int64, index int64) (int, error) {
	if index < 0 || index > 63 {
		return 0, errors.New("Invalid index. Index must be in the range [0, 63].")
	}

	var mask int64 = 1 << index
	if (num & mask) != 0 {
		return 1, nil
	}
	return 0, nil
	// return (num & mask) != 0, nil
}

func printBits(num int64) {
	for i := int64(63); i >= 0; i-- {
		tmp, _ := getBit(num, i)
		fmt.Printf("%d", tmp)
	}
	fmt.Println()
}

func main() {
	var num int64 = -2
	printBits(num)
	var num1 int64 = -9223372036854775808 
	printBits(num1)
	var num2 int64 = 9223372036854775807
	printBits(num2)
}
