package main


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
