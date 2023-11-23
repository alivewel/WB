package main

import "fmt"

type Cat struct {
	sound string
}

func newCat() *Cat {
	cat := &Cat{sound: "Meow!"}
	return cat
}

func (a *Cat) makeSound() {
	fmt.Println(a.sound)
}
