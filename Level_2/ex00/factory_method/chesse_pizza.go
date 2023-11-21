package main

type CheesePizza struct {
	name string
}

func (p *CheesePizza) GetPizzaName() string {
	return p.name
}
