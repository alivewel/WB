package main

import (
	"fmt"
	"math"

	"github.com/shopspring/decimal"
)

// Разработать программу, которая перемножает, делит, складывает, вычитает
// две числовых переменных a,b, значение которых > 2^20.

func main() {
	int1 := int64(math.Pow(2, 30))
	int2 := int64(math.Pow(2, 20))

	num1 := decimal.NewFromInt(int1)
	num2 := decimal.NewFromInt(int2)

	sum := num1.Add(num2)
	subtraction := num1.Sub(num2)
	product := num1.Mul(num2)
	division, _ := num1.Div(num2).Round(2).Float64()

	fmt.Println("Сумма:", sum)
	fmt.Println("Разность:", subtraction)
	fmt.Println("Произведение:", product)
	fmt.Println("Деление:", division)
}
