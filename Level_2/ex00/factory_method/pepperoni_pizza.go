package main

type PepperoniPizza struct {
	name string
}

func (p *PepperoniPizza) GetPizzaName() string {
	return p.name
}
