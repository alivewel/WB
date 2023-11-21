package main

import (
	"fmt"
)

type Cook struct{}

func (c *Cook) MakeBurger() {
	fmt.Println("Make Burger!")
}

func (c *Cook) MakeFries() {
	fmt.Println("Make Fries!")
}
