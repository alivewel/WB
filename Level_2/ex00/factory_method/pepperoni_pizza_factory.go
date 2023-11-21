package main

type PepperoniPizzaFactory struct{}

func (f PepperoniPizzaFactory) CreatePizza() Pizza {
	return &PepperoniPizza{name: "Pepperoni Pizza"}
}
