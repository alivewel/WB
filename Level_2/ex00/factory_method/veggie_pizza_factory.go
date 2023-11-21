package main

type VeggiePizzaFactory struct{}

func (f VeggiePizzaFactory) CreatePizza() Pizza {
	return &CheesePizza{name: "Veggie Pizza"}
}
