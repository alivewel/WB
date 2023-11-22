package main

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
