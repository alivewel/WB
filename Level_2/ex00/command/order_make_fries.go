package main

type OrderMakeFries struct {
	cook *Cook
}

func (c *OrderMakeFries) Execute() {
	c.cook.MakeFries()
}
