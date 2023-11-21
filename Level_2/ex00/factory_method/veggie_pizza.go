package main

type VeggiePizza struct {
	name string
}

func (p *VeggiePizza) GetPizzaName() string {
	return p.name
}
