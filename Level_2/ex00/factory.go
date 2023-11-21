package main

import "fmt"

type Pizza interface {
	GetPizzaName() string
}

type CheesePizza struct {
	name string
}

func (p *CheesePizza) GetPizzaName() string {
	return p.name
}

type PepperoniPizza struct {
	name string
}

func (p *PepperoniPizza) GetPizzaName() string {
	return p.name
}

type VeggiePizza struct {
	name string
}

func (p *VeggiePizza) GetPizzaName() string {
	return p.name
}

type PizzaFactory interface {
	CreatePizza() Pizza
}

type CheesePizzaFactory struct{}

func (f CheesePizzaFactory) CreatePizza() Pizza {
	return &CheesePizza{name: "Cheese Pizza"}
}

type PepperoniPizzaFactory struct{}

func (f PepperoniPizzaFactory) CreatePizza() Pizza {
	return &PepperoniPizza{name: "Pepperoni Pizza"}
}

type VeggiePizzaFactory struct{}

func (f VeggiePizzaFactory) CreatePizza() Pizza {
	return &CheesePizza{name: "Veggie Pizza"}
}

func main() {
	CheesePizzaFactory := CheesePizzaFactory{}
	pizza := CheesePizzaFactory.CreatePizza()
	fmt.Println(pizza.GetPizzaName())

	PepperoniPizzaFactory := PepperoniPizzaFactory{}
	pizza = PepperoniPizzaFactory.CreatePizza()
	fmt.Println(pizza.GetPizzaName())

	VeggiePizzaFactory := VeggiePizzaFactory{}
	pizza = VeggiePizzaFactory.CreatePizza()
	fmt.Println(pizza.GetPizzaName())
}
