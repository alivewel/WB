package main

type Context struct {
	animal Animal
}

func (c *Context) SetAnimal(a Animal) {
	c.animal = a
}

func (c *Context) MakeAnimalSound() {
	c.animal.makeSound()
}
