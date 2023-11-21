package main

import "fmt"

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
