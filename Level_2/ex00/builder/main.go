package main

import "fmt"

type House struct {
	TypeHouse  string
	Color      string
	CountFloor int
}

type Builder interface {
	SetTypeHouse()
	SetColor()
	SetCountFloor()
	BuildHouse() House
}

type WoodHouseBuilder struct {
	house House
}

func (w *WoodHouseBuilder) SetTypeHouse() {
	w.house.TypeHouse = "wood"
}

func (w *WoodHouseBuilder) SetColor() {
	w.house.Color = "brown"
}

func (w *WoodHouseBuilder) SetCountFloor() {
	w.house.CountFloor = 1
}

func (w *WoodHouseBuilder) BuildHouse() House {
	return House{
		TypeHouse:  w.house.TypeHouse,
		Color:      w.house.Color,
		CountFloor: w.house.CountFloor,
	}
}

type BrickHouseBuilder struct {
	house House
}

func (b *BrickHouseBuilder) SetTypeHouse() {
	b.house.TypeHouse = "brick"
}

func (b *BrickHouseBuilder) SetColor() {
	b.house.Color = "red"
}

func (b *BrickHouseBuilder) SetCountFloor() {
	b.house.CountFloor = 3
}

func (b *BrickHouseBuilder) BuildHouse() House {
	return House{
		TypeHouse:  b.house.TypeHouse,
		Color:      b.house.Color,
		CountFloor: b.house.CountFloor,
	}
}

type Director struct {
	builder Builder
}

func NerDirector(b Builder) *Director {
	return &Director{builder: b}
}

func (d *Director) SetBuilder(b Builder) {
	d.builder = b
}

func (d *Director) BuildHouse() House {
	d.builder.SetTypeHouse()
	d.builder.SetColor()
	d.builder.SetCountFloor()
	return d.builder.BuildHouse()
}

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
