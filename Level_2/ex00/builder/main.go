package main

import "fmt"

func main() {
	woodHouseBuilder := WoodHouseBuilder{}
	brickHouseBuilder := BrickHouseBuilder{}

	director := NerDirector(&woodHouseBuilder)
	woodHouse := director.BuildHouse()
	fmt.Println("Wood house:", woodHouse)

	director.SetBuilder(&brickHouseBuilder)
	brickHouse := director.BuildHouse()
	fmt.Println("Brick house:", brickHouse)
}

// Команда для запуска:
// go run .
