package main

type CheesePizzaFactory struct{}

func (f CheesePizzaFactory) CreatePizza() Pizza {
	return &CheesePizza{name: "Cheese Pizza"}
}
