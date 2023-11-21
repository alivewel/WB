package main

type OrderMakeBurger struct {
	cook *Cook
}

func (c *OrderMakeBurger) Execute() {
	c.cook.MakeBurger()
}
