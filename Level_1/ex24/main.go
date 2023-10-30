package main

import (
	"fmt"
	"math"
)

// Разработать программу нахождения расстояния между двумя точками,
// которые представлены в виде структуры Point с инкапсулированными
// параметрами x,y и конструктором.

type Point struct {
	x int
	y int
}

func NewPoint(x int, y int) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func GetDistance(point1, point2 *Point) float64 {
	num1 := math.Pow(float64(point2.x-point1.x), 2)
	num2 := math.Pow(float64(point2.y-point1.y), 2)
	result := math.Sqrt(num1 + num2)
	return result
}

func main() {
	point1 := NewPoint(11, 2)
	point2 := NewPoint(22, 7)

	result := GetDistance(point1, point2)
	fmt.Printf("result %.2f", result)
}
