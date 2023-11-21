package main

import (
	"fmt"
)

type Order interface {
	Execute()
}

type OrderMakeBurger struct {
	cook *Cook
}

func (c *OrderMakeBurger) Execute() {
	c.cook.MakeBurger()
}

type OrderMakeFries struct {
	cook *Cook
}

func (c *OrderMakeFries) Execute() {
	c.cook.MakeFries()
}

type Cook struct{}

func (c *Cook) MakeBurger() {
	fmt.Println("Make Burger!")
}

func (c *Cook) MakeFries() {
	fmt.Println("Make Fries!")
}

type Waitress struct{
	order Order
}

func (w *Waitress) SetOrder(order Order) {
	w.order = order
}

func (w *Waitress) ExecuteOrder() {
	w.order.Execute()
}

func main() {
	cook := &Cook{}
	
	makeBurger := &OrderMakeBurger{cook: cook}
	makeFries := &OrderMakeFries{cook: cook}
	
	waitress := &Waitress{}

	waitress.SetOrder(makeBurger)
	waitress.ExecuteOrder()

	waitress.SetOrder(makeFries)
	waitress.ExecuteOrder()
}
