package main

import (
	"errors"
	"fmt"
)

// первый бит имеет индекс 0
// последний бит имеет индекс 63

func getBit(num int64, index int) (int, error) {
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

func setBit(num *int64, index int, bit int) error {
	if bit != 1 && bit != 0 {
		return errors.New("Invalid bit value. Bit must be 0 or 1.")
	}
	if index < 0 || index > 63 {
		return errors.New("Invalid index. Index must be in the range [0, 63].")
	}
	if bit == 1 {
		*num = *num | 1<<index
	} else {
		*num = *num &^ (1 << index)
	}
	return nil
}

func printBits(num int64) {
	for i := 63; i >= 0; i-- {
		tmp, _ := getBit(num, i)
		fmt.Printf("%d", tmp)
	}
	fmt.Println()
}

func main() {
	// var num int64 = 0
	var num int64 = 9223372036854775807
	printBits(num)

	// for i := 0; i <= 63; i++ {
	// 	err := setBit(&num, i, 0)
	// 	if err != nil {
	// 		fmt.Println("Ошибка:", err)
	// 	}
	// 	printBits(num)
	// }
	var num1 int64 = 9223372036854775807
	err := setBit(&num1, 63, 1)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
	printBits(num1)
	print(num1)

	// var num1 int64 = -9223372036854775808
	// printBits(num1)
	// var num2 int64 = 9223372036854775807
	// printBits(num2)
}
