package main

import (
	"fmt"
)

type Dog struct {
	sound string
}

func newDog() *Dog {
	dog := &Dog{sound: "Bark!"}
	return dog
}

func (a *Dog) makeSound() {
	fmt.Println(a.sound)
}
