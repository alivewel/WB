package main

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
